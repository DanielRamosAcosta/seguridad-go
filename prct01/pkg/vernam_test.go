package prct01_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/danielramosacosta/seguridad/prct01/pkg"
)

var _ = Describe("vernam", func() {
	Describe("CipherMessage", func() {
		It("ciphers a message", func() {
			const originalMessage = "SOL"
			const randomKey = "001111000001100001110011"

			cipheredMessage := CipherMessage(originalMessage, randomKey)
			Expect(cipheredMessage).To(Equal("oW?"))
		})
	})
	Describe("DecipherMessage", func() {
		It("deciphers a message", func() {
			const cipheredMessage = "[t"
			const randomKey = "0000111100100001"

			originalMessage := DecipherMessage(cipheredMessage, randomKey)
			Expect(originalMessage).To(Equal("TU"))
		})
	})
})
