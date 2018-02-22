package prct03

import (
	"fmt"
	"log"
)

// RC4 creates a pseudo random generator using the RC4 stream cipher algorithm
type RC4 struct {
	S []byte
	i uint
	f uint
}

// New creates a new generator using the given seed
func New(seed []byte) RC4 {
	g := RC4{S: []byte{}, i: 0, f: 0}
	g.S = g.KeySchedulingAlgorithm(seed)
	return g
}

// KeySchedulingAlgorithm initializates the S array
func (g *RC4) KeySchedulingAlgorithm(seed []byte) []byte {
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

// PseudoRandomGenerationAlgorithm generates a new random byte
func (g *RC4) PseudoRandomGenerationAlgorithm() byte {
	g.i = (g.i + 1) % 256
	g.f = (g.f + uint(g.S[g.i])) % 256

	g.S[g.i], g.S[g.f] = g.S[g.f], g.S[g.i]
	t := (uint(g.S[g.i]) + uint(g.S[g.f])) % 256

	format := "Byte %d de secuencia cifrante: Salida = S[%d]=%d \t\t%08b"
	log.Printf(format, g.i, t, g.S[t], g.S[t])

	return g.S[t]
}
