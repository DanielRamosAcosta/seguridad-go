package rc4

import (
	"fmt"
	"log"
)

func KeySchedulingAlgorithm(seed []byte) []byte {
	var f uint
	var S []byte
	var K []byte

	for i := 0; i <= 255; i++ {
		S = append(S, byte(i))
		K = append(K, seed[i%len(seed)])
	}

	log.Println("S=", S)
	log.Println("K=", K)

	for i := 0; i <= 255; i++ {
		f = (f + uint(S[i]) + uint(K[i])) % 256
		S[i], S[f] = S[f], S[i]
		log.Println(
			fmt.Sprintf(
				"S[%d]=%d, K[%d]=%d produce f=%d, S[%d]=%d, S[%d]=%d",
				i, S[f], i, K[i], f, i, S[i], f, S[f],
			),
		)
	}

	return S
}

func PseudoRandomGenerationAlgorithm(_i byte, _f byte, _S []byte) (byte, byte, byte, []byte) {
	S := make([]byte, len(_S))
	copy(S, _S)

	var i = (uint(_i) + 1) % 256
	var f = (uint(_f) + uint(S[i])) % 256

	S[i], S[f] = S[f], S[i]
	t := (uint(S[i]) + uint(S[f])) % 256

	log.Printf("Byte %d de secuencia cifrante: Salida = S[%d]=%d \t\t%08b", i, t, S[t], S[t])

	return S[t], byte(i), byte(f), S
}
