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
	"strings"
)

type Type struct {
	protoType      *ProtoType
	Name           string // the name of the API call, Event or Type
	Description    string // the description/comments for the struct
	GoType         string // the type this would be in Go
	UnderlyingType string // the type defined in protocol.json
	EnumVals       string // if it's an enum string list out the possible values as a comment
	IsSubType      bool   // is this a sub type? (Should be prefixed with Sub in template)
	Properties     []*TypeProperties
}

func NewType(protoType *ProtoType) *Type {
	t := &Type{}
	t.protoType = protoType
	t.Name = protoType.Id
	t.Description = protoType.Description
	if t.Description == "" {
		t.Description = "No Description."
	}
	t.UnderlyingType = protoType.Type
	t.Properties = make([]*TypeProperties, 0)
	return t
}

func NewSubType(parentProps *TypeProperties, protoProps *ProtoProperty) *Type {
	st := &Type{}
	st.IsSubType = true
	// only convert to type if it's an array
	if protoProps.IsArray() {
		st.protoType = typeFromProperties(protoProps)
	}

	st.Name = "Sub" + strings.Title(parentProps.Name)
	st.Properties = make([]*TypeProperties, 0)
	st.UnderlyingType = parentProps.UnderlyingType
	return st
}

func (t *Type) IsNonPropertiesObject() bool {
	return (t.UnderlyingType == "object" && len(t.protoType.Properties) == 0)
}

func (t *Type) GetUnderlyingType() string {
	return t.UnderlyingType
}

func (t *Type) IsArray() bool {
	return t.UnderlyingType == "array"
}

func (t *Type) GetArrayType() string {
	if t.protoType.Items.Type != "" {
		return t.protoType.Items.Type
	}

	if t.protoType.Items.Ref != "" {
		return t.protoType.Items.Ref
	}
	return "object"
}
