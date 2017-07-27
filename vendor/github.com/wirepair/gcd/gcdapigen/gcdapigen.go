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
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"
)

type GlobalReference struct {
	LocalRefName    string // the local reference name "StackTrace"
	ExternalGoName  string // the gotype name "ConsoleStackTrace"
	IsBaseType      bool   // is this a base type (int, string, float64 etc)
	BaseType        string // what is it
	EnumDescription string // is it an enum, if so keep metadata of enum values
	IsArrayRef      bool   // is this a reference to another type that is of an array type
}

// Stores a list of all references and if they are base types
var globalRefs map[string]*GlobalReference

const (
	debug               = false
	outputDir           = "output"
	prefix              = "chrome_"
	templateFile        = "api_template.txt"
	browserProtocolFile = "https://chromium.googlesource.com/chromium/src/+/master/third_party/WebKit/Source/core/inspector/browser_protocol.json?format=text"
	jsProtocolFile      = "https://chromium.googlesource.com/v8/v8/+/master/src/inspector/js_protocol.json?format=text"
)

var file string
var update bool
var templates *template.Template // for code generation
var funcMap template.FuncMap     // helper funcs

func init() {
	flag.BoolVar(&update, "update", false, "download and merge js_protocol.json and browser_protocol.json into protocol.json")
	flag.StringVar(&file, "file", "protocol.json", "open remote debugger protocol file, if -update the filename to write to.")
	funcMap := template.FuncMap{
		"Title":    strings.Title,
		"ToLower":  strings.ToLower,
		"Reserved": modifyReserved,
		"NullType": nullType,
	}
	templates = template.Must(template.New(templateFile).Funcs(funcMap).ParseFiles(templateFile))
}

func main() {
	var domains []*Domain
	globalRefs = make(map[string]*GlobalReference)

	flag.Parse()

	if update {
		download(browserProtocolFile, jsProtocolFile)
		return
	}

	protocolData := openFile()
	if debug == false {
		createOutputDirectory()
	}

	protocolApi := unmarshalProtocol(protocolData)
	major := protocolApi.Version.Major
	minor := protocolApi.Version.Minor

	// iterate over the protocol once to resolve references
	for _, proto := range protocolApi.Domains {
		PopulateReferences(proto.Domain, proto.Types)
	}

	for _, protoDomain := range protocolApi.Domains {
		domain := NewDomain(major, minor, protoDomain.Domain)
		fmt.Printf("Creating api for domain: %s\n", protoDomain.Domain)

		// Do types first
		if protoDomain.Types != nil && len(protoDomain.Types) > 0 {
			domain.PopulateTypes(protoDomain.Types)
		}

		// Next Events
		if protoDomain.Events != nil && len(protoDomain.Events) > 0 {
			domain.PopulateEvents(protoDomain.Events)
		}

		// Then Commands
		if protoDomain.Commands != nil && len(protoDomain.Commands) > 0 {
			domain.PopulateCommands(protoDomain.Commands)
		}
		domains = append(domains, domain)
		domain.ResolveImports()
		domain.WriteDomain()
	}

	//for k, v := range globalRefs {
	//	fmt.Printf("ref: %s value: %#v\n", k, v)
	//}
}

func openFile() []byte {
	log.Printf("Opening %s for reading...\n", file)
	protocolData, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("Error reading file: %v\n", err)
	}
	return protocolData
}

func createOutputDirectory() {

	if _, err := os.Stat(outputDir); os.IsExist(err) {

		return
	}

	err := os.Mkdir(outputDir, 0755)

	if err != nil && os.IsExist(err) {
		log.Printf("output directory %s already exists\n", outputDir)
		return
	} else if err != nil {
		log.Fatalf("error creating api output directory")
	}
}

func unmarshalProtocol(protocolData []byte) *ProtoDebuggerApi {
	api := &ProtoDebuggerApi{}
	err := json.Unmarshal(protocolData, api)
	if err != nil {
		switch u := err.(type) {
		case *json.SyntaxError:
			log.Fatalf("syntax error at offset %d: %s\n", u.Offset, u)
		case *json.UnmarshalTypeError:
			log.Fatal("unmarshal type err: " + err.Error() + " " + err.(*json.UnmarshalTypeError).Value)
		case *json.InvalidUnmarshalError:
			log.Fatal("InvalidUnmarshalError: " + err.Error())
		default:
			log.Fatal("UnmarshalError: " + err.Error())
		}
	}
	return api
}
