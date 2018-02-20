package main

import (
	"fmt"
)

func main() {
	const originalMessage = "SOL"
	const randomKey = "001111000001100001110011"
	fmt.Println("Entrada:")
	fmt.Println("Mensaje original:", originalMessage)
	messageBytes := ascciStrToBytes(originalMessage)
	randomKeyBytes := binaryStrToBytes(randomKey)
	fmt.Println("Mensaje original en binario:", bytesToBinaryStr(messageBytes))
	fmt.Println("Entrada:")
	fmt.Println("Clave aleatoria:", randomKey)

	// ascciStrToBytes works
	expectedMessageBytes := []byte{0x53, 0x4F, 0x4C}
	for i, n := range expectedMessageBytes {
		if messageBytes[i] != n {
			panic("No coincinden los messageBytes")
		}
	}

	// binaryStrToBytes works
	expectedRandomKeyBytes := []byte{0x3C, 0x18, 0x73}
	for i, n := range expectedRandomKeyBytes {
		if randomKeyBytes[i] != n {
			panic("No coincinden los randomKeyBytes")
		}
	}

	// bytesToBinaryStr works
	const expectedBinaryMessageString = "010100110100111101001100"
	res := bytesToBinaryStr(messageBytes)
	if expectedBinaryMessageString != res {
		fmt.Println(expectedBinaryMessageString)
		fmt.Println(res)
		panic("No son iguales!")
	}

	// xorBytesArray works
	aaaaexpectedMessageBytes := []byte{0x53, 0x4F, 0x4C}
	aaaaexpectedRandomKeyBytes := []byte{0x3C, 0x18, 0x73}
	resss := xorBytesArray(aaaaexpectedMessageBytes, aaaaexpectedRandomKeyBytes)
	expectedXOR := []byte{0x6F, 0x57, 0x3F}

	for i, n := range expectedXOR {
		if resss[i] != n {
			fmt.Println(i)
			fmt.Println(fmt.Sprintf("%x", n))
			fmt.Println(fmt.Sprintf("%x", resss[i]))
			panic("No coincinden los xor")
		}
	}

	cifredoBinaryStr := bytesToBinaryStr(resss)

	// bytesToBinaryStr works
	const expectedMessageCifredBinary = "011011110101011100111111"
	if expectedMessageCifredBinary != cifredoBinaryStr {
		fmt.Println(expectedMessageCifredBinary)
		fmt.Println(cifredoBinaryStr)
		panic("No son iguales!")
	}

	mensajeCifrado := bytesToAscii(resss)

	const expectedMensajeCifrado = "oW?"
	if expectedMensajeCifrado != mensajeCifrado {
		fmt.Println(expectedMensajeCifrado)
		fmt.Println(mensajeCifrado)
		panic("No son iguales!")
	}

	fmt.Println("Mensaje cifrado en binario:", cifredoBinaryStr)
	fmt.Println("Mensaje cifrado:", mensajeCifrado)
}

/*

* Entrada:
* Mensaje original: SOL
* Salida:
* Mensaje original en binario: 010100110100111101001100
* Longitud del mensaje binario: 24
* Entrada:
* Clave aleatoria: 001111000001100001110011
* Salida:
* Mensaje cifrado en binario: 011011110101011100111111
* Mensaje cifrado: oW?
 */
