package main

import (
	"log"

	"github.com/danielramosacosta/seguridad/prct03/rc4"
)

func main() {
	seed := []byte{0x02, 0x05}
	message := []byte{0x01, 0x22}

	log.Println("Inicialización:")
	S := rc4.KeySchedulingAlgorithm(seed)

	i := byte(0)
	f := byte(0)
	var generatedByte byte

	for index, c := range message {
		generatedByte, i, f, S = rc4.PseudoRandomGenerationAlgorithm(i, f, S)
		cipheredText := c ^ generatedByte

		log.Printf("Byte %d de texto original: Entrada: M[%d]=%d \t\t\t%08b", index+1, index+1, c, c)
		log.Printf("Byte %d de texto cifrado: Salida: C[%d]=%d\t\t\t%08b", index+1, index+1, cipheredText, cipheredText)
	}
}
