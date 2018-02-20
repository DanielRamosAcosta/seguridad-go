package crypto

import (
	"log"
	"github.com/danielramosacosta/seguridad/prct01/utils"
)

/**
 * Converts an ASCII string to a Byte array.
 * Example: "SOL", "001111000001100001110011" => "oW?"
 */
func CipherMessage(messageStr string, randomKeyBinaryStr string) string {
	messageBytes := utils.AscciStrToBytes(messageStr)
	randomKeyBytes := utils.BinaryStrToBytes(randomKeyBinaryStr)

	log.Println("Entrada:")
	log.Println("Mensaje original:", messageStr)
	log.Println("Mensaje original en binario:", utils.BytesToBinaryStr(messageBytes))
	log.Println("Entrada:")
	log.Println("Clave aleatoria:", randomKeyBinaryStr)

	cipheredMessageBytes := utils.XorBytesArray(messageBytes, randomKeyBytes)
	cipheredMessageStr := utils.BytesToAsciiStr(cipheredMessageBytes)

	log.Println("Mensaje cifrado en binario:", utils.BytesToBinaryStr(cipheredMessageBytes))
	log.Println("Mensaje cifrado:", cipheredMessageStr)
	return cipheredMessageStr
}

/**
 * Converts an ASCII string to a Byte array.
 * Example: "[t", "0000111100100001" => "TU"
 */
func DecipherMessage(cipheredMessage string, randomKeyBinaryStr string) string {
	cipheredMessageBytes := utils.AscciStrToBytes(cipheredMessage)
	randomKeyBytes := utils.BinaryStrToBytes(randomKeyBinaryStr)

	log.Println("Entrada:")
	log.Println("Mensaje cifrado:", cipheredMessage)
	log.Println("Salida:")
	log.Println("Mensaje cifrado en binario:", utils.BytesToBinaryStr(cipheredMessageBytes))
	log.Println("Longitud del mensaje binario:", len(cipheredMessageBytes)*8)
	log.Println("Entrada:")
	log.Println("Clave aleatoria:", randomKeyBinaryStr)

	originalMessageBytes := utils.XorBytesArray(cipheredMessageBytes, randomKeyBytes)
	originalMessageStr := utils.BytesToAsciiStr(originalMessageBytes)

	log.Println("Salida:")
	log.Println("Mensaje original en binario:", utils.BytesToBinaryStr(originalMessageBytes))
	log.Println("Mensaje original en binario:", originalMessageStr)

	return originalMessageStr
}
