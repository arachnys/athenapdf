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

type Return struct {
	protoReturn    *ProtoCommandReturns
	Name           string // property name
	Description    string // property description
	UnderlyingType string
	GoType         string
	Optional       bool   // is this property optional?
	EnumVals       string // possible enum values as a string
	IsRef          bool   // is a reference to another type
	IsTypeArray    bool   // for templates to spit out []
	IsPointer      bool
	Ref            string
}

func NewReturn(protoReturn *ProtoCommandReturns) *Return {
	r := &Return{}
	r.protoReturn = protoReturn
	r.Name = protoReturn.Name
	r.Description = protoReturn.Description
	r.UnderlyingType = protoReturn.Type
	r.Ref = protoReturn.Ref

	if protoReturn.Ref != "" {
		r.IsRef = true
	}

	// if array, check underlying array type to see if it's a reference
	if r.IsArray() {
		r.IsTypeArray = true
		if arrayRef := r.ArrayRef(); arrayRef != "" {
			r.Ref = arrayRef
			r.IsRef = true
		}
	}

	return r
}

func (r *Return) ArrayRef() string {
	if r.protoReturn.Items.Ref != "" {
		return r.protoReturn.Items.Ref
	}
	return ""
}

// PropSetter interface methods
func (r *Return) GetGoType() string {
	return r.GoType
}

func (r *Return) SetIsTypeArray(isTypeArray bool) {
	r.IsTypeArray = true
}

func (r *Return) GetEnumVals() string {
	return r.EnumVals
}

func (r *Return) GetRef() string {
	return r.Ref
}

func (r *Return) SetIsRef(isRef bool) {
	r.IsRef = isRef
}

func (r *Return) GetIsRef() bool {
	return r.IsRef
}

func (r *Return) SetGoType(goType string) {
	r.GoType = goType
}

func (r *Return) GetDescription() string {
	return r.Description
}
func (r *Return) SetDescription(description string) {
	r.Description = description
}

func (r *Return) SetPointerType(isPointer bool) {
	r.IsPointer = isPointer
}

// SharedProperties interface methods
func (r *Return) IsNonPropertiesObject() bool {
	return (r.UnderlyingType == "object") // return with object type never has properties // && len(r.protoReturn.Properties) == 0)
}

func (r *Return) GetUnderlyingType() string {
	return r.UnderlyingType
}

func (r *Return) IsArray() bool {
	return r.UnderlyingType == "array"
}

func (r *Return) GetArrayType() string {
	if r.protoReturn.Items.Type != "" {
		return r.protoReturn.Items.Type
	}

	if r.protoReturn.Items.Ref != "" {
		return r.protoReturn.Items.Ref
	}
	return "object"
}
