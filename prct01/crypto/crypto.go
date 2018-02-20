package crypto

import (
	"fmt"

	"github.com/danielramosacosta/seguridad/prct01/utils"
)

/**
 * Converts an ASCII string to a Byte array.
 * Example: "SOL", "001111000001100001110011" => "oW?"
 */
func CipherMessage(messageStr string, randomKeyBinaryStr string) string {
	messageBytes := utils.AscciStrToBytes(messageStr)
	randomKeyBytes := utils.BinaryStrToBytes(randomKeyBinaryStr)

	fmt.Println("Entrada:")
	fmt.Println("Mensaje original:", messageStr)
	fmt.Println("Mensaje original en binario:", utils.BytesToBinaryStr(messageBytes))
	fmt.Println("Entrada:")
	fmt.Println("Clave aleatoria:", randomKeyBinaryStr)

	cipheredMessageBytes := utils.XorBytesArray(messageBytes, randomKeyBytes)
	cipheredMessageStr := utils.BytesToAsciiStr(cipheredMessageBytes)

	fmt.Println("Mensaje cifrado en binario:", utils.BytesToBinaryStr(cipheredMessageBytes))
	fmt.Println("Mensaje cifrado:", cipheredMessageStr)
	return cipheredMessageStr
}

/**
 * Converts an ASCII string to a Byte array.
 * Example: "[t", "0000111100100001" => "TU"
 */
func DecipherMessage(cipheredMessage string, randomKeyBinaryStr string) string {
	cipheredMessageBytes := utils.AscciStrToBytes(cipheredMessage)
	randomKeyBytes := utils.BinaryStrToBytes(randomKeyBinaryStr)

	fmt.Println("Entrada:")
	fmt.Println("Mensaje cifrado:", cipheredMessage)
	fmt.Println("Salida:")
	fmt.Println("Mensaje cifrado en binario:", utils.BytesToBinaryStr(cipheredMessageBytes))
	fmt.Println("Longitud del mensaje binario:", len(cipheredMessageBytes)*8)
	fmt.Println("Entrada:")
	fmt.Println("Clave aleatoria:", randomKeyBinaryStr)

	originalMessageBytes := utils.XorBytesArray(cipheredMessageBytes, randomKeyBytes)
	originalMessageStr := utils.BytesToAsciiStr(originalMessageBytes)

	fmt.Println("Salida:")
	fmt.Println("Mensaje original en binario:", utils.BytesToBinaryStr(originalMessageBytes))
	fmt.Println("Mensaje original en binario:", originalMessageStr)

	return originalMessageStr
}
