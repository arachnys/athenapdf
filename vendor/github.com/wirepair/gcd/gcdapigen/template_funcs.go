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

func modifyReserved(input string) string {
	switch input {
	case "type":
		return "theType"
	case "range":
		return "theRange"
	case "interface":
		return "theInterface"
	case "for":
		return "theFor"
	}
	return input
}

func nullType(input string) string {
	if strings.Contains(input, "[]") {
		return "nil"
	}
	//fmt.Printf("INPUT: %s\n", input)
	switch input {
	case "int":
		return "0"
	case "float64":
		return "0"
	case "string":
		return "\"\""
	case "bool":
		return "false"
	case "interface{}":
		return "nil"
	}
	return "nil"
}
