package main

import (
	"log"

	"github.com/danielramosacosta/seguridad/prct01/crypto"
)

func main() {
	crypto.CipherMessage("SOL", "001111000001100001110011")
	log.Println(" ------- ")
	crypto.DecipherMessage("[t", "0000111100100001")
}
