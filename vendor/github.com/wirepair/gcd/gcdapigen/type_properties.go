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

// Used for Types and Parameters to function calls
type TypeProperties struct {
	protoProperty  *ProtoProperty
	Name           string // property name
	Description    string // property description
	UnderlyingType string
	GoType         string
	Optional       bool   // is this property optional?
	EnumVals       string // possible enum values as a string
	Ref            string
	IsRef          bool // is a reference to another type
	IsPointer      bool // should we output as pointer (for API types, not basic types)
	IsTypeArray    bool // for templates to spit out []
}

func NewTypeProperties(props *ProtoProperty) *TypeProperties {
	tp := &TypeProperties{}
	tp.protoProperty = props
	tp.Name = props.Name
	tp.Description = props.Description
	tp.Optional = props.Optional
	tp.Ref = props.Ref
	tp.UnderlyingType = props.Type
	if tp.IsArray() {
		tp.IsTypeArray = true
		if arrayRef := tp.ArrayRef(); arrayRef != "" {
			tp.Ref = arrayRef
			tp.IsRef = true
		}
	}
	return tp
}

// PropSetter interface methods
func (p *TypeProperties) GetGoType() string {
	return p.GoType
}

func (p *TypeProperties) SetGoType(goType string) {
	p.GoType = goType
}

func (p *TypeProperties) GetIsRef() bool {
	return p.IsRef
}

func (p *TypeProperties) SetIsRef(isRef bool) {
	p.IsRef = isRef
}

func (p *TypeProperties) SetIsTypeArray(isTypeArray bool) {
	p.IsTypeArray = true
}

func (p *TypeProperties) GetRef() string {
	return p.Ref
}

func (p *TypeProperties) GetEnumVals() string {
	return p.EnumVals
}

func (p *TypeProperties) GetDescription() string {
	return p.Description
}
func (p *TypeProperties) SetDescription(description string) {
	p.Description = description
}

func (p *TypeProperties) SetPointerType(isPointer bool) {
	p.IsPointer = isPointer
}

// SharedProperties interface methods
func (p *TypeProperties) IsNonPropertiesObject() bool {
	return (p.UnderlyingType == "object" && len(p.protoProperty.Properties) == 0)
}

func (p *TypeProperties) GetUnderlyingType() string {
	return p.UnderlyingType
}

func (p *TypeProperties) IsArray() bool {
	return p.UnderlyingType == "array"
}

func (p *TypeProperties) GetArrayType() string {
	if p.protoProperty.Items.Type != "" {
		return p.protoProperty.Items.Type
	}

	if p.protoProperty.Items.Ref != "" {
		return p.protoProperty.Items.Ref
	}
	return "object"
}

func (p *TypeProperties) ArrayRef() string {
	if p.protoProperty.Items.Ref != "" {
		return p.protoProperty.Items.Ref
	}
	return ""
}
