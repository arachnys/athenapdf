/*
The MIT License (MIT)

Copyright (c) 2016 isaac dawson

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

type Domain struct {
	Major    string // major api version
	Minor    string // minor api version
	Filename string
	Domain   string
	Imports  []string
	Hidden   bool
	SubTypes []*Type
	Types    []*Type
	Events   []*Event
	Commands []*Command
	// basicTypes holds a map of type.RefName and type.Underlying type so we can replace $ref
	// with the underlying type (provided it's not another object or array)

	//typeMap map[string]*BaseType
}

func NewDomain(major, minor, domain string) *Domain {
	d := &Domain{Major: major, Minor: minor, Domain: domain}
	d.Types = make([]*Type, 0)
	d.SubTypes = make([]*Type, 0)
	d.Events = make([]*Event, 0)
	d.Commands = make([]*Command, 0)
	//d.typeMap = make(map[string]*BaseType)
	return d
}

// Extract each type and call handleType, add the result to our Types slice.
func (d *Domain) PopulateTypes(types []*ProtoType) {
	// do first pass to get all underlying type information
	for _, protoType := range types {
		fmt.Printf("Populating type: %s\n", protoType.Id)
		newType := NewType(protoType)
		// igore empty property types as we turn those into Refs
		if len(protoType.Properties) > 0 {
			d.handleType(newType, protoType.Properties)
			d.Types = append(d.Types, newType)
		}

	}
}

func (d *Domain) PopulateEvents(events []*ProtoEvent) {
	for _, protoEvent := range events {
		newEvent := NewEvent(protoEvent)

		if newEvent.HasParams {
			d.handleEvents(newEvent, protoEvent.Parameters)
			d.Events = append(d.Events, newEvent)
		} else {
			fmt.Printf("event: %s has no params\n", newEvent.Name)
			fmt.Printf("protoEvent: %#v\n", protoEvent)
		}
		// don't add structs that don't have parameters.
	}
}

func (d *Domain) PopulateCommands(commands []*ProtoCommand) {
	for _, protoCommand := range commands {
		newCmd := NewCommand(protoCommand)

		if newCmd.HasParams {
			d.handleParameters(newCmd, protoCommand.Parameters)
		}

		if newCmd.HasReturn {
			d.handleReturns(newCmd, protoCommand.Returns)
		}
		d.Commands = append(d.Commands, newCmd)
	}
}

func (d *Domain) handleEvents(newEvent *Event, protoParameters []*ProtoProperty) {
	for _, protoParam := range protoParameters {
		newParam := NewTypeProperties(protoParam)
		if isBaseType(newParam) {
			goType := getGoType(newParam)
			d.createBase(newParam, goType)
			newEvent.Parameters = append(newEvent.Parameters, newParam)
			continue
		}

		if ok := d.resolveReference(newParam); ok {
			newEvent.Parameters = append(newEvent.Parameters, newParam)
		}
	}
}

func (d *Domain) handleParameters(newCmd *Command, protoParameters []*ProtoProperty) {
	for _, protoParam := range protoParameters {
		newParam := NewTypeProperties(protoParam)
		if isBaseType(newParam) {
			goType := getGoType(newParam)
			d.createBase(newParam, goType)
			newCmd.Parameters = append(newCmd.Parameters, newParam)
			continue
		}

		if ok := d.resolveReference(newParam); ok {
			newCmd.Parameters = append(newCmd.Parameters, newParam)
		}
	}
}

func (d *Domain) handleReturns(newCmd *Command, protoReturns []*ProtoCommandReturns) {
	for _, protoReturn := range protoReturns {
		newReturn := NewReturn(protoReturn)
		if isBaseType(newReturn) {
			goType := getGoType(newReturn)
			d.createBase(newReturn, goType)
			newCmd.Returns = append(newCmd.Returns, newReturn)
			continue
		}

		if ok := d.resolveReference(newReturn); ok {
			newCmd.Returns = append(newCmd.Returns, newReturn)
		}
	}
}

// Takes in a new type, checks if it is a base type, or an object or an array.
func (d *Domain) handleType(newType *Type, typeProperties []*ProtoProperty) {
	// loop over properties of this new type
	for _, protoProp := range typeProperties {
		// It's a reference, see if it points to a base type or not
		newProp := NewTypeProperties(protoProp)
		// base type, add it and fix up description
		if isBaseType(protoProp) {
			goType := getGoType(newProp)
			d.createBase(newProp, goType)
			newType.Properties = append(newType.Properties, newProp)
			continue
		}

		if ok := d.resolveReference(newProp); ok {
			newType.Properties = append(newType.Properties, newProp)
			continue
		}

		// is this a subType?
		if isSubType(protoProp) {
			fmt.Printf("issubtype true for: %s\n", protoProp.Name)
			d.createSubType(newType, protoProp)
		}
	}
}

func (d *Domain) createBase(prop PropSetter, goType string) {
	prop.SetGoType(goType)
	if len(prop.GetEnumVals()) > 0 {
		prop.SetDescription(prop.GetDescription() + "Enum values: " + prop.GetEnumVals())
	}

	if isPointerType(prop) {
		prop.SetPointerType(true)
	}

	return
}

// Creates a new SubType *Type object. This is for nested structs that are better
// defined outside of the original Type. It will call handleType to iterate over
// the nested properties to create it in much of the same way as a normal type
// except we prefix it with the Sub keyword.
func (d *Domain) createSubType(newType *Type, protoProp *ProtoProperty) {
	// create the property
	newProp := NewTypeProperties(protoProp)
	// create the new sub type
	subType := NewSubType(newProp, protoProp)

	// recursive to add props to this type
	recursiveProps := protoProp.Properties
	// array type requires the protoType properties.
	if subType.IsArray() {
		recursiveProps = subType.protoType.Properties
	}
	d.handleType(subType, recursiveProps)
	d.SubTypes = append(d.SubTypes, subType)
	// update ref from property to new sub type
	refName := d.Domain + "Sub" + strings.Title(protoProp.Name)
	newProp.GoType = refName
	newProp.IsRef = true
	newType.Properties = append(newType.Properties, newProp)
}

// Since we don't want useless single underlying type struct definitions everywhere we resolve
// any references to single underlying type objects.
func (d *Domain) resolveReference(prop PropSetter) bool {

	if prop.GetRef() == "" {
		fmt.Printf("ref was empty: %s\n", prop.GetGoType())
		return false
	}
	refName := ""
	// Local reference, add . between domain/ref type so we can look it up in our globalRefs
	if !strings.Contains(prop.GetRef(), ".") {
		refName = d.Domain + "." + prop.GetRef()
	} else {
		refName = prop.GetRef()
	}

	ref := globalRefs[refName]
	fmt.Printf("REF (%s): %#v\n", refName, ref)
	// base type
	if ref.IsBaseType {
		prop.SetGoType(ref.BaseType)
	} else {
		prop.SetGoType(ref.ExternalGoName)
		prop.SetIsRef(true)
	}
	// set the type as being an array of whatever type it references
	if ref.IsArrayRef {
		prop.SetIsTypeArray(true)
	}
	// add enum possible values to description
	if ref.EnumDescription != "" {
		prop.SetDescription(prop.GetDescription() + ref.EnumDescription)
	}

	if isPointerType(prop) {
		prop.SetPointerType(true)
	}
	return true
}

// if any command has a return statement we use json, so import it.
func (d *Domain) ResolveImports() {
	for _, returns := range d.Commands {
		if returns.HasReturn {
			d.Imports = append(d.Imports, "encoding/json")
			return
		}
	}
}

func (d *Domain) WriteDomain() {
	if debug {
		wr := os.Stdout

		err := templates.ExecuteTemplate(wr, templateFile, d)
		if err != nil {
			log.Fatalf("Error writing to template: %s\n", err)
		}
		return
	}

	domainFile := outputDir + string(os.PathSeparator) + strings.ToLower(d.Domain) + ".go"
	wr, err := os.Create(domainFile)
	if err != nil {
		log.Fatalf("error creating output file: %s\n", err)
	}

	err = templates.ExecuteTemplate(wr, templateFile, d)
	if err != nil {
		log.Fatalf("Error writing to template: %s\n", err)
	}

	// clean up file
	cmd := exec.Command("gofmt", "-w", domainFile)
	var out bytes.Buffer
	var outErr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &outErr
	err = cmd.Start()
	if err != nil {
		log.Fatalf("error running gofmt! %s\n", err)
	}

	err = cmd.Wait()
	if err != nil {
		fmt.Printf("error in domain: %s\n%s", d.Domain, outErr.String())
		log.Fatalf("error waiting for gofmt to complete: %s\n", err)
	}
}
