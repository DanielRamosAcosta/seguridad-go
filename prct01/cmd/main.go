package main

import (
	"log"

	"github.com/danielramosacosta/seguridad/prct01/pkg"
)

func main() {
	prct01.CipherMessage("SOL", "001111000001100001110011")
	log.Println(" ------- ")
	prct01.DecipherMessage("[t", "0000111100100001")
}
