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

	log.Println("Inicialización:")
	log.Println("S=", S)
	log.Println("K=", K)

	for i := 0; i < 256; i++ {
		f = (f + uint(S[i]) + uint(K[i])) % 256
		S[i], S[f] = S[f], S[i]
		log.Println(
			fmt.Sprintf(
				"S[%d]=%d, K[%d]=%d produce f=%d, S[%d]=%d, S[%d]=%d",
				i, S[i], i, K[i], f, i, S[f], f, S[i],
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

	return S[t], byte(i), byte(f), S
}
