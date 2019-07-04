package protoc_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestProtoc(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Protoc Suite")
}
