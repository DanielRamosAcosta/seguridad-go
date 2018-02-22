package prct02_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/danielramosacosta/seguridad/prct02/pkg"
)

var _ = Describe("utils", func() {
	Describe("CleanString", func() {
		It("cleans the given string", func() {
			const str = "-this IS NOT cleanñ-"
			const expectedStr = "THISISNOTCLEAN"

			resString := prct02.CleanString(str)

			Expect(resString).To(Equal(expectedStr))
		})
	})
})
