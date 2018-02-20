package utils_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/danielramosacosta/seguridad/prct01/utils"
)

var _ = Describe("Utils", func() {
	Describe("AscciStrToBytes", func() {
		It("converts an ASCII string into a byte array", func() {
			const str = "SOL"
			expectedBytes := []byte{0x53, 0x4F, 0x4C}
			resultBytes := utils.AscciStrToBytes(str)
			Expect(resultBytes).To(Equal(expectedBytes))
		})
	})
	Describe("BinaryStrToBytes", func() {
		It("converts a binary string into a byte array", func() {
			const str = "010111000100111110011010"
			expectedBytes := []byte{0x5C, 0x4F, 0x9A}
			resultBytes := utils.BinaryStrToBytes(str)

			Expect(resultBytes).To(Equal(expectedBytes))
		})
	})
	Describe("BytesToBinaryStr", func() {
		It("converts byte array into a binary string", func() {
			bytes := []byte{0x5C, 0x4F, 0x9A}
			const expectedStr = "010111000100111110011010"
			resStr := utils.BytesToBinaryStr(bytes)

			Expect(resStr).To(Equal(expectedStr))
		})
	})
	Describe("BytesToAsciiStr", func() {
		It("converts byte array into its ascii representation", func() {
			bytes := []byte{0x53, 0x4F, 0x4C}
			const expectedStr = "SOL"
			resStr := utils.BytesToAsciiStr(bytes)

			Expect(resStr).To(Equal(expectedStr))
		})
	})
	Describe("AscciStrToBinaryStr", func() {
		It("converts an ascii string into its binary representation", func() {
			const str = "SOL"
			const expectedStr = "010100110100111101001100"
			resStr := utils.AscciStrToBinaryStr(str)

			Expect(resStr).To(Equal(expectedStr))
		})
	})
	Describe("XorBytesArray", func() {
		It("does XOR operation between two byte array", func() {
			bytes1 := []byte{0x53, 0x4F, 0x4C}
			bytes2 := []byte{0x3C, 0x18, 0x73}
			expectedBytes := []byte{0x6F, 0x57, 0x3F}
			resultBytes := utils.XorBytesArray(bytes1, bytes2)

			Expect(resultBytes).To(Equal(expectedBytes))
		})
	})
})
