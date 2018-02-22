package prct01

import (
	"log"
)

// CipherMessage ciphers the given string using Vernam's cipher
// Example: "SOL", "001111000001100001110011" => "oW?"
func CipherMessage(messageStr string, randomKeyBinaryStr string) string {
	messageBytes := AscciStrToBytes(messageStr)
	randomKeyBytes := BinaryStrToBytes(randomKeyBinaryStr)

	log.Println("Entrada:")
	log.Println("Mensaje original:", messageStr)
	log.Println("Mensaje original en binario:", BytesToBinaryStr(messageBytes))
	log.Println("Entrada:")
	log.Println("Clave aleatoria:", randomKeyBinaryStr)

	cipheredMessageBytes := XorBytesArray(messageBytes, randomKeyBytes)
	cipheredMessageStr := BytesToASCIIStr(cipheredMessageBytes)

	log.Println("Mensaje cifrado en binario:", BytesToBinaryStr(cipheredMessageBytes))
	log.Println("Mensaje cifrado:", cipheredMessageStr)
	return cipheredMessageStr
}

// DecipherMessage deciphers the given string using Vernam's cipher
// Example: "[t", "0000111100100001" => "TU"
func DecipherMessage(cipheredMessage string, randomKeyBinaryStr string) string {
	cipheredMessageBytes := AscciStrToBytes(cipheredMessage)
	randomKeyBytes := BinaryStrToBytes(randomKeyBinaryStr)

	log.Println("Entrada:")
	log.Println("Mensaje cifrado:", cipheredMessage)
	log.Println("Salida:")
	log.Println("Mensaje cifrado en binario:", BytesToBinaryStr(cipheredMessageBytes))
	log.Println("Longitud del mensaje binario:", len(cipheredMessageBytes)*8)
	log.Println("Entrada:")
	log.Println("Clave aleatoria:", randomKeyBinaryStr)

	originalMessageBytes := XorBytesArray(cipheredMessageBytes, randomKeyBytes)
	originalMessageStr := BytesToASCIIStr(originalMessageBytes)

	log.Println("Salida:")
	log.Println("Mensaje original en binario:", BytesToBinaryStr(originalMessageBytes))
	log.Println("Mensaje original en binario:", originalMessageStr)

	return originalMessageStr
}
