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

type Command struct {
	Name               string
	Description        string
	Parameters         []*TypeProperties
	Returns            []*Return
	HasParams          bool
	HasReturn          bool
	NoParamReturnCalls bool
	ParamCalls         bool
	ReturnCalls        bool
	ParamReturnCalls   bool
}

func NewCommand(protoCommand *ProtoCommand) *Command {
	c := &Command{}
	c.Name = protoCommand.Name
	c.Description = protoCommand.Description
	if protoCommand.Parameters != nil && len(protoCommand.Parameters) > 0 {
		c.HasParams = true
	}

	if protoCommand.Returns != nil && len(protoCommand.Returns) > 0 {
		c.HasReturn = true
	}
	// Determine type of call for template output
	if c.HasParams == false && c.HasReturn == false {
		c.NoParamReturnCalls = true
	}

	if c.HasParams == true && c.HasReturn == false {
		c.ParamCalls = true
	}

	if c.HasParams == false && c.HasReturn == true {
		c.ReturnCalls = true
	}

	if c.HasParams == true && c.HasReturn == true {
		c.ParamReturnCalls = true
	}

	c.Returns = make([]*Return, 0)
	c.Parameters = make([]*TypeProperties, 0)
	return c
}
