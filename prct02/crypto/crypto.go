package crypto

import (
	"bytes"

	"github.com/danielramosacosta/seguridad/prct02/utils"
)

/**
 * Ciphers the given string using Vigenere's cipher
 * Example: "ESTE MENSAJE SE AUTODESTRUIRA", "MISION" => "QALMARZASRSFQIMBCQQALZIVDI"
 */
func CipherMessage(originalMessage string, key string) string {
	var buffer bytes.Buffer

	for i, c := range utils.CleanString(originalMessage) {
		keyIndex := i % len(key)
		keyChar := key[keyIndex]
		originalChar := byte(c)

		finalChar := ((originalChar + keyChar) % 26) + 0x41

		buffer.WriteString(string(finalChar))
	}

	return buffer.String()
}

/**
 * Deciphers the given string using Vigenere's cipher
 * Example: "QALMARZASRSFQIMBCQQALZIVDI", "MISION" => "ESTEMENSAJESEAUTODESTRUIRA"
 */
func DecipherMessage(cipheredMessage string, key string) string {
	var buffer bytes.Buffer

	for i, c := range utils.CleanString(cipheredMessage) {
		keyIndex := i % len(key)
		keyChar := key[keyIndex]
		originalChar := byte(c)

		finalChar := ((originalChar - keyChar + 26) % 26) + 0x41

		buffer.WriteString(string(finalChar))
	}

	return buffer.String()
}
