package main

import (
	"log"

	"github.com/danielramosacosta/seguridad/prct02/crypto"
)

func main() {
	const message = "ESTE MENSAJE SE AUTODESTRURIA"
	const key = "MISION"
	ciphered := crypto.CipherMessage(message, key)

	log.Println("Palabra clave:", key)
	log.Println("Texto original:", message)
	log.Println("Texto cifrado:", ciphered)
}
