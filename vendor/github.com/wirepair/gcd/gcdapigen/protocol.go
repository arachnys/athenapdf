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

// Top level API protocol file
type ProtoDebuggerApi struct {
	Version *ProtoApiVersion `json:"version"`
	Domains []*ProtoDomain   `json:"domains"`
}

// Version information
type ProtoApiVersion struct {
	Major string `json:"major"`
	Minor string `json:"minor"`
}

// The Domain (contains all objects, their type/commands/events)
type ProtoDomain struct {
	Domain      string          `json:"domain"`
	Description string          `json:"description,omitempty"`
	Types       []*ProtoType    `json:"types,omitempty"`
	Commands    []*ProtoCommand `json:"commands,omitempty"`
	Events      []*ProtoEvent   `json:"events,omitempty"`
	Hidden      bool            `json:"hidden,omitempty"`
	Items       *ProtoItem      `json:"items,omitempty"`
}

// A Type which represents objects specific to the API method
type ProtoType struct {
	Id          string           `json:"id"`
	Type        string           `json:"type"`
	Description string           `json:"description,omitempty"`
	Enum        []string         `json:"enum,omitempty"`
	Properties  []*ProtoProperty `json:"properties,omitempty"`
	Hidden      bool             `json:"hidden,omitempty"`
	Items       *ProtoItem       `json:"items,omitempty"`
	MinItems    int64            `json:"minItems,omitempty"`
	MaxItems    int64            `json:"maxItems,omitempty"`
}

func (p *ProtoType) IsNonPropertiesObject() bool {
	return (p.Type == "object" && len(p.Properties) == 0)
}

func (p *ProtoType) GetUnderlyingType() string {
	return p.Type
}

func (p *ProtoType) GetArrayType() string {
	if p.Type != "array" || p.Items == nil {
		return ""
	}
	if p.Items.Type != "" {
		return p.Items.Type
	}
	if p.Items.Ref != "" {
		return p.Items.Ref
	}
	return ""
}

func (p *ProtoType) IsArray() bool {
	return p.Type == "array"
}

// A property & Parameter type used by both commands & types
type ProtoProperty struct {
	Name        string           `json:"name"`
	Type        string           `json:"type,omitempty"`
	Description string           `json:"description,omitempty"`
	Ref         string           `json:"$ref,omitempty"`
	Optional    bool             `json:"optional,omitempty"`
	Hidden      bool             `json:"hidden,omitempty"`
	Enum        []string         `json:"enum,omitempty"`
	Items       *ProtoItem       `json:"items,omitempty"`
	Properties  []*ProtoProperty `json:"properties,omitempty"`
}

func (p *ProtoProperty) IsNonPropertiesObject() bool {
	return (p.Type == "object" && len(p.Properties) == 0)
}

func (p *ProtoProperty) GetUnderlyingType() string {
	return p.Type
}

func (p *ProtoProperty) GetArrayType() string {
	if p.Type != "array" || p.Items == nil {
		return ""
	}
	if p.Items.Type != "" {
		return p.Items.Type
	}
	if p.Items.Ref != "" {
		return p.Items.Ref
	}
	return ""
}

func (p *ProtoProperty) IsArray() bool {
	return p.Type == "array"
}

// An item used by types, properties and events.
type ProtoItem struct {
	Type        string           `json:"type,omitempty"`
	Ref         string           `json:"$ref,omitempty"`
	Properties  []*ProtoProperty `json:"properties,omitempty"`
	Description string           `json:"description,omitempty"`
	Enum        []string         `json:"enum,omitempty"`
}

// The API Command call.
type ProtoCommand struct {
	Name        string                 `json:"name"`
	Type        string                 `json:"type,omitempty"`
	Description string                 `json:"description,omitempty"`
	Handlers    []string               `json:"handlers,omitempty"`
	Parameters  []*ProtoProperty       `json:"parameters,omitempty"`
	Returns     []*ProtoCommandReturns `json:"returns,omitempty"`
	Hidden      bool                   `json:"hidden,omitempty"`
	Async       bool                   `json:"async,omitempty"`
	Redirect    string                 `json:"redirect,omitempty"`
}

// The return parameters for an API call
type ProtoCommandReturns struct {
	Name        string     `json:"name"`
	Type        string     `json:"type,omitempty"`
	Ref         string     `json:"$ref,omitempty"`
	Items       *ProtoItem `json:"items,omitempty"`
	Description string     `json:"description,omitempty"`
}

// An event, asynchronous events that can come in once
// enabled.
type ProtoEvent struct {
	Name        string           `json:"name"`
	Type        string           `json:"type,omitempty"`
	Description string           `json:"description,omitempty"`
	Ref         string           `json:"$ref,omitempty"`
	Optional    bool             `json:"optional,omitempty"`
	Hidden      bool             `json:"hidden,omitempty"`
	Enum        []string         `json:"enum,omitempty"`
	Items       *ProtoItem       `json:"items,omitempty"`
	Parameters  []*ProtoProperty `json:"parameters,omitempty"`
}

func (p *ProtoEvent) IsNonPropertiesObject() bool {
	return (p.Type == "object" && len(p.Parameters) == 0)
}

func (p *ProtoEvent) GetUnderlyingType() string {
	return p.Type
}

func (p *ProtoEvent) GetArrayType() string {
	if p.Type != "array" || p.Items == nil {
		return ""
	}
	if p.Items.Type != "" {
		return p.Items.Type
	}
	if p.Items.Ref != "" {
		return p.Items.Ref
	}
	return ""
}

func (p *ProtoEvent) IsArray() bool {
	return p.Type == "array"
}
