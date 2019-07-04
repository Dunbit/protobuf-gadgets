package version_test

import (
	"regexp"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/dunbit/protobuf-gadgets/pkg/version"
)

var _ = Describe("Version", func() {

	Describe("AppVersion", func() {

		It("Should have a valid version format", func() {
			match, err := regexp.MatchString("^v(\\d+)\\.(\\d+)\\.(\\d+)$", AppVersion)

			Expect(err).NotTo(HaveOccurred())
			Expect(match).To(BeTrue())
		})
	})
})
