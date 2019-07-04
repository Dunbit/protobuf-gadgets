package protoc_test

import (
	"bytes"

	plugin_go "github.com/golang/protobuf/protoc-gen-go/plugin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/dunbit/protobuf-gadgets/pkg/protoc"
)

type dummyPlugin struct{}

func (p *dummyPlugin) Generate(req *plugin_go.CodeGeneratorRequest) (*plugin_go.CodeGeneratorResponse, error) {
	return nil, nil
}

var _ = Describe("Plugin", func() {

	Describe("RunWithIO", func() {

		It("TODO Should add a test", func() {
			var input = []byte("")
			var output = []byte("")

			reader := bytes.NewBuffer(input)
			writer := bytes.NewBuffer(output)

			plugin := new(dummyPlugin)
			RunWithIO(plugin, reader, writer)

			Expect(true).To(BeTrue())
		})
	})
})
