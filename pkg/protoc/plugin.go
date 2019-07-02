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

package protoc

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/gogo/protobuf/proto"
	plugin_go "github.com/golang/protobuf/protoc-gen-go/plugin"
)

// Plugin defines a protoc plugin
type Plugin interface {
	Generate(req *plugin_go.CodeGeneratorRequest) (*plugin_go.CodeGeneratorResponse, error)
}

// Run ...
func Run(p Plugin) {
	RunWithIO(p, os.Stdin, os.Stdout)
}

// RunWithIO ...
func RunWithIO(p Plugin, reader io.Reader, writer io.Writer) {
	req, err := readIO(reader)
	if err != nil {
		writeError(writer, err)
		return
	}

	resp, err := p.Generate(req)
	if err != nil {
		writeError(writer, err)
		return
	}

	err = writeIO(writer, resp)
	if err != nil {
		writeError(writer, err)
	}
}

// readIO ...
func readIO(reader io.Reader) (*plugin_go.CodeGeneratorRequest, error) {
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		os.Exit(1)
	}

	req := new(plugin_go.CodeGeneratorRequest)
	err = proto.Unmarshal(data, req)
	if err != nil {
		os.Exit(1)
	}

	if len(req.GetFileToGenerate()) == 0 {
		return nil, fmt.Errorf("no files to generate")
	}

	return req, nil
}

// writeIO ...
func writeIO(writer io.Writer, resp *plugin_go.CodeGeneratorResponse) error {
	data, err := proto.Marshal(resp)
	if err != nil {
		return err
	}

	_, err = writer.Write(data)
	if err != nil {
		return err
	}

	return nil
}

// writeError ...
func writeError(writer io.Writer, err error) {
	s := err.Error()

	resp := &plugin_go.CodeGeneratorResponse{
		Error: &s,
	}

	data, err := proto.Marshal(resp)
	if err != nil {
		os.Exit(1)
	}

	_, err = writer.Write(data)
	if err != nil {
		os.Exit(1)
	}
}
