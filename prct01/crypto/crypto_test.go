package crypto_test

import (
	"testing"

	"github.com/danielramosacosta/seguridad/prct01/crypto"
)

func TestCipherMessage(t *testing.T) {
	const message = "SOL"
	const randomKey = "001111000001100001110011"
	const expetedOutput = "oW?"

	cipheredMessage := crypto.CipherMessage(message, randomKey)

	if cipheredMessage != expetedOutput {
		t.Errorf("Results doesn't match %s != %s", expetedOutput, cipheredMessage)
	}
}

func TestDecipherMessage(t *testing.T) {
	const cipheredMessage = "[t"
	const randomKey = "0000111100100001"
	const expectedMessage = "TU"

	resultMessage := crypto.DecipherMessage(cipheredMessage, randomKey)

	if resultMessage != expectedMessage {
		t.Errorf("Results doesn't match %s != %s", resultMessage, expectedMessage)
	}
}
