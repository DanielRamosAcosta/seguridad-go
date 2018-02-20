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

	bytes := []byte("")

	for i := 0; i < len(str); i += 8 {
		substr := str[i : i+8]

		if i, err := strconv.ParseInt(substr, 2, 64); err != nil {
			panic(err)
		} else {
			bytes = append(bytes, byte(i))
		}
	}
	return bytes
}

/**
 * Converts a Byte array to a binary string.
 * Example: [ 0x5C, 0x4F, 0x9A ] => "010111000100111110011010"
 */
func BytesToBinaryStr(bytes []byte) string {
	var binaryStr = ""
	for _, n := range bytes {
		var rawBinaryStr = fmt.Sprintf("%b", n)

		for len(rawBinaryStr) < 8 {
			rawBinaryStr = "0" + rawBinaryStr
		}

		binaryStr = binaryStr + rawBinaryStr
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
func XorBytesArray(byt1 []byte, byt2 []byte) []byte {
	bytes := []byte("")
	for i, n := range byt1 {
		bytes = append(bytes, n^byt2[i])
	}
	return bytes
}
