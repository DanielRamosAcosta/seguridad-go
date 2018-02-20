package utils_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/danielramosacosta/seguridad/prct02/utils"
)

var _ = Describe("Utils", func() {
	Describe("CleanString", func() {
		It("cleans the given string", func() {
			const str = "-this IS NOT cleanñ-"
			const expectedStr = "THISISNOTCLEAN"

			resString := utils.CleanString(str)

			Expect(resString).To(Equal(expectedStr))
		})
	})
})
