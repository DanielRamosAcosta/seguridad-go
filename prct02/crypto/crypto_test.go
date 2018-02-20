package crypto_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/danielramosacosta/seguridad/prct02/crypto"
)

var _ = Describe("crypto", func() {
	Describe("CipherMessage", func() {
		It("ciphers a message", func() {
			const originalMessage = "ESTE MENSAJE SE AUTODESTRUIRA"
			const key = "MISION"

			cipheredMessage := crypto.CipherMessage(originalMessage, key)
			Expect(cipheredMessage).To(Equal("QALMARZASRSFQIMBCQQALZIVDI"))
		})
	})
	Describe("DecipherMessage", func() {
		It("deciphers a message", func() {
			const cipheredMessage = "QALMARZASRSFQIMBCQQALZIVDI"
			const key = "MISION"

			originalMessage := crypto.DecipherMessage(cipheredMessage, key)
			Expect(originalMessage).To(Equal("ESTEMENSAJESEAUTODESTRUIRA"))
		})
	})
})
