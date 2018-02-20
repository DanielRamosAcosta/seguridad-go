package utils_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/danielramosacosta/seguridad/prct01/utils"
)

var _ = Describe("Utils", func() {
	Describe("AscciStrToBytes", func() {
		It("Debe poder convertir un string a un array de bytes", func() {
			const str = "SOL"
			expectedBytes := []byte{0x53, 0x4F, 0x4C}
			resultBytes := utils.AscciStrToBytes(str)
			Expect(expectedBytes).To(Equal(resultBytes))
		})
	})
})
