// Copyright 2019 Authors of protobuf-gadgets
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package docmaster

import (
	"fmt"
	"strings"

	"github.com/dunbit/protobuf-gadgets/pkg/protoc"
	plugin_go "github.com/golang/protobuf/protoc-gen-go/plugin"
)

// Format to render
type Format int

// Format values
const (
	NoFormat Format = 0
	Markdown Format = 1
	HTML     Format = 2
)

// Options ...
type Options struct {
	Format       Format
	OutputFile   string
	TemplateFile string
}

// ParseOptions ...
// example: format=template,template=index.txt,file=docs.html
func ParseOptions(req *plugin_go.CodeGeneratorRequest) (*Options, error) {
	param := req.GetParameter()

	o := &Options{}

	if param == "" {
		return o, nil
	}

	elements := strings.Split(param, ",")
	for _, element := range elements {
		opt := strings.Split(element, "=")
		if len(opt) != 2 {
			return nil, fmt.Errorf("Invalid parameters: %s", param)
		}
		switch opt[0] {
		case "format":
			switch opt[1] {
			case "markdown":
				o.Format = Markdown
			case "html":
				o.Format = HTML
			}
		case "file":
			o.OutputFile = opt[1]
		case "template":
			o.TemplateFile = opt[1]
		}
	}

	return o, nil
}

// Plugin defines a protoc plugin
type Plugin struct {
}

// NewPlugin ...
func NewPlugin() *Plugin {
	return &Plugin{}
}

// Generate ...
func (d *Plugin) Generate(req *plugin_go.CodeGeneratorRequest) (*plugin_go.CodeGeneratorResponse, error) {
	options, err := ParseOptions(req)
	if err != nil {
		return nil, err
	}

	data, err := protoc.ParseCodeGeneratorRequest(req)
	if err != nil {
		return nil, err
	}

	file, err := RenderTemplates(options, data)
	if err != nil {
		return nil, err
	}

	return &plugin_go.CodeGeneratorResponse{
		File: file,
	}, nil
}
