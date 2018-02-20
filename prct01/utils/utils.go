package utils

import (
	"fmt"
	"strconv"
)

/**
 * Converts an ASCII string to a Byte array.
 * Example: "SOL" => [ 0x53, 0x4F, 0x4C ]
 */
func AscciStrToBytes(str string) []byte {
	return []byte(str)
}

/**
 * Converts an Binary string to a Byte array.
 * Example: "010111000100111110011010" => [ 0x5C, 0x4F, 0x9A ]
 */
func BinaryStrToBytes(str string) []byte {
	if (len(str) % 8) != 0 {
		panic("Malformed binary string")
	}

	numberBytes := len(str) / 8
	bytes := make([]byte, numberBytes)

	for i := 0; i < numberBytes; i++ {
		startIndex := i * 8
		endIndex := (i + 1) * 8
		substr := str[startIndex:endIndex]
		numberByte, _ := strconv.ParseInt(substr, 2, 64)
		bytes[i] = byte(numberByte)
	}
	return bytes
}

/**
 * Converts a Byte array to a binary string.
 * Example: [ 0x5C, 0x4F, 0x9A ] => "010111000100111110011010"
 */
func BytesToBinaryStr(bytes []byte) string {
	var binaryStr string
	for _, n := range bytes {
		binaryStr = binaryStr + fmt.Sprintf("%08b", n)
	}
	return binaryStr
}

/**
 * Converts a Byte array to its corresponding ASCII string
 * Example: [ 0x53, 0x4F, 0x4C ] => "SOL"
 */
func BytesToAsciiStr(bytes []byte) string {
	return string(bytes[:len(bytes)])
}

/**
 * Converts ASCII string to a Binary string.
 * Example: "SOL" => "010100110100111101001100"
 */
func AscciStrToBinaryStr(str string) string {
	bytes := AscciStrToBytes(str)
	return BytesToBinaryStr(bytes)
}

/**
 * Does the XOR operator between two byte array
 * Example: [0x53, 0x4F, 0x4C] ^ [0x3C, 0x18, 0x73] =  [0x6F, 0x57, 0x3F]
 */
func XorBytesArray(bytes1 []byte, bytes2 []byte) []byte {
	if len(bytes1) != len(bytes2) {
		panic("Byte array length mismatch")
	}

	var bytes = make([]byte, len(bytes1))
	for i := range bytes1 {
		byte1 := bytes1[i]
		byte2 := bytes2[i]
		bytes[i] = byte1 ^ byte2
	}
	return bytes
}
