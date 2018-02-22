package prct01

import (
	"bytes"
	"fmt"
	"strconv"
)

// AscciStrToBytes converts an ASCII string to a Byte array.
// Example: "SOL" => [ 0x53, 0x4F, 0x4C ]
func AscciStrToBytes(str string) []byte {
	return []byte(str)
}

// BinaryStrToBytes converts an Binary string to a Byte array.
// Example: "010111000100111110011010" => [ 0x5C, 0x4F, 0x9A ]
func BinaryStrToBytes(str string) []byte {
	if (len(str) % 8) != 0 {
		panic("Malformed binary string")
	}

	var buffer bytes.Buffer

	for i := 0; i < len(str); i += 8 {
		substr := str[i : i+8]
		numberByte, _ := strconv.ParseInt(substr, 2, 64)
		buffer.WriteByte(byte(numberByte))
	}
	return buffer.Bytes()
}

// BytesToBinaryStr converts a Byte array to a binary string.
// Example: [ 0x5C, 0x4F, 0x9A ] => "010111000100111110011010
func BytesToBinaryStr(bytesArr []byte) string {
	var buffer bytes.Buffer
	for _, n := range bytesArr {
		buffer.WriteString(fmt.Sprintf("%08b", n))
	}
	return buffer.String()
}

// BytesToASCIIStr converts a Byte array to its corresponding ASCII string
// Example: [ 0x53, 0x4F, 0x4C ] => "SOL"
func BytesToASCIIStr(bytes []byte) string {
	return string(bytes[:len(bytes)])
}

// AscciStrToBinaryStr converts ASCII string to a Binary string.
// Example: "SOL" => "010100110100111101001100"
func AscciStrToBinaryStr(str string) string {
	bytes := AscciStrToBytes(str)
	return BytesToBinaryStr(bytes)
}

// XorBytesArray does the XOR operator between two byte array
// Example: [0x53, 0x4F, 0x4C] ^ [0x3C, 0x18, 0x73] =  [0x6F, 0x57, 0x3F]
func XorBytesArray(bytes1 []byte, bytes2 []byte) []byte {
	if len(bytes1) != len(bytes2) {
		panic("Byte array length mismatch")
	}

	var buffer bytes.Buffer
	for i := range bytes1 {
		byte1 := bytes1[i]
		byte2 := bytes2[i]
		buffer.WriteByte(byte1 ^ byte2)
	}
	return buffer.Bytes()
}
