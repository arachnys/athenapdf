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
	"fmt"
	"strings"
)

type PropSetter interface {
	GetRef() string
	GetIsRef() bool
	GetGoType() string
	GetEnumVals() string
	SetGoType(goType string)
	SetIsTypeArray(isTypeArray bool)
	GetDescription() string
	SetDescription(description string)
	SetIsRef(isRef bool)
	SetPointerType(isPointer bool)
}

type SharedProperties interface {
	GetUnderlyingType() string
	GetArrayType() string
	IsArray() bool
	IsNonPropertiesObject() bool
}

// Loops over the domain ProtoType slice to create a reference map that we can use
// to resolve references when we go to spit out templates. We check if the type is a base
// type like:
// {
//     "id": "RequestId",
//     "type": "string",
//     "description": "Unique request identifier."
// },
// Instead of spitting out "type NetworkRequestId string" and having other structs
// include references to it, we resolve any references directly to 'string'
// this makes working with the API far easier (no stupid type conversions everywhere)
func PopulateReferences(domain string, types []*ProtoType) {
	for _, protoType := range types {
		ref := &GlobalReference{}
		ref.LocalRefName = protoType.Id
		if isBaseType(protoType) {
			ref.IsBaseType = true
			ref.BaseType = getGoType(protoType)
		}

		ref.ExternalGoName = domain + protoType.Id
		// an array pointing to another type
		if protoType.IsArray() {
			arrType := protoType.GetArrayType()
			ref.IsArrayRef = true
			ref.ExternalGoName = domain + arrType
		}

		if len(protoType.Enum) > 0 {
			ref.EnumDescription = " enum values: " + strings.Join(protoType.Enum, ", ")
		}
		globalRefs[domain+"."+protoType.Id] = ref

		if ref.IsBaseType == false {
			populateSubReferences(domain, protoType)
		}
	}
}

// Check if the type has a nested type and create a new sub type for it and add to our
// reference map.
func populateSubReferences(domain string, protoType *ProtoType) {
	for _, prop := range protoType.Properties {
		// a new SubType
		if isSubType(prop) {
			prefix := "Sub" + strings.Title(prop.Name)
			ref := &GlobalReference{}
			ref.LocalRefName = prop.Name
			ref.ExternalGoName = domain + prefix
			globalRefs[domain+prefix] = ref
			continue
		}

	}
}

// Get the go type, if it is an array we look up the underlying type of the array.
func getGoType(props SharedProperties) string {

	protoType := props.GetUnderlyingType()
	if props.IsArray() {
		protoType = props.GetArrayType()
	}

	returnType := ""
	switch protoType {
	case "string", "enum":
		returnType = "string"
	case "integer":
		returnType = "int"
	case "number":
		returnType = "float64"
	case "object":
		returnType = "map[string]interface{}" // default
	case "any":
		returnType = "interface{}"
	case "boolean":
		returnType = "bool"
	}
	return returnType
}

// Is this a base type? If NoPropertiesObject we just set it to map[string]interface{}
// If it's an array, we look up if the array underlying type is a base type.
func isBaseType(props SharedProperties) bool {
	if props.IsNonPropertiesObject() {
		return true
	}

	underlyingType := props.GetUnderlyingType()
	if props.IsArray() {
		underlyingType = props.GetArrayType()
	}

	switch underlyingType {
	case "any", "string", "enum", "integer", "number", "boolean":
		return true
	}
	return false
}

// Do the properties of this type point to a nested type? Can be array or object that nests
// a sub type.
func isSubType(protoProp *ProtoProperty) bool {
	if (protoProp.Type == "object" && !protoProp.IsNonPropertiesObject()) || (protoProp.Type == "array" && protoProp.Items.Type == "object") {
		return true
	}
	return false
}

func isPointerType(props PropSetter) bool {
	goType := props.GetGoType()
	switch goType {
	case "int", "string", "bool", "float64", "map[string]interface{}", "interface{}":
		return false
	}
	return true
}

// used for nested arrays, convert properties to a type
func typeFromProperties(protoProps *ProtoProperty) *ProtoType {
	t := &ProtoType{}
	t.Type = protoProps.Type
	t.Enum = protoProps.Enum
	t.Hidden = protoProps.Hidden
	t.Properties = protoProps.Items.Properties
	t.Id = protoProps.Name
	t.Description = protoProps.Description
	fmt.Printf("RETURNING NEW TYPE FROM PROPS: %#v\n", t)
	return t
}
