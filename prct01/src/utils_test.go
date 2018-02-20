package utils_test

import (
	"testing"

	"github.com/danielramosacosta/seguridad/prct01/src"
)

func TestAscciStrToBytes(t *testing.T) {
	const str = "SOL"
	expectedBytes := []byte{0x53, 0x4F, 0x4C}
	resultBytes := utils.AscciStrToBytes(str)

	for i, expectedByte := range expectedBytes {
		resultByte := resultBytes[i]
		if resultByte != expectedByte {
			t.Errorf("Bytes doesn't match 0x%x != 0x%x", expectedByte, resultByte)
		}
	}
}

func TestBinaryStrToBytes(t *testing.T) {
	const str = "010111000100111110011010"
	expectedBytes := []byte{0x5C, 0x4F, 0x9A}
	resultBytes := utils.BinaryStrToBytes(str)

	for i, expectedByte := range expectedBytes {
		resultByte := resultBytes[i]
		if resultByte != expectedByte {
			t.Errorf("Bytes doesn't match 0x%x != 0x%x", expectedByte, resultByte)
		}
	}
}

func TestBytesToBinaryStr(t *testing.T) {
	bytes := []byte{0x5C, 0x4F, 0x9A}
	const expectedStr = "010111000100111110011010"
	resStr := utils.BytesToBinaryStr(bytes)

	if expectedStr != resStr {
		t.Errorf("Strings doesn't match %s != 0x%s", expectedStr, resStr)
	}
}

func TestBytesToAscii(t *testing.T) {
	bytes := []byte{0x53, 0x4F, 0x4C}
	const expectedStr = "SOL"
	resStr := utils.BytesToAscii(bytes)

	if expectedStr != resStr {
		t.Errorf("Strings doesn't match %s != 0x%s", expectedStr, resStr)
	}
}

func TestAscciStrToBinaryStr(t *testing.T) {
	const str = "SOL"
	const expectedStr = "010100110100111101001100"

	resStr := utils.AscciStrToBinaryStr(str)

	if expectedStr != resStr {
		t.Errorf("Strings doesn't match %s != 0x%s", expectedStr, resStr)
	}
}

func TestXorBytesArray(t *testing.T) {
	bytes1 := []byte{0x53, 0x4F, 0x4C}
	bytes2 := []byte{0x3C, 0x18, 0x73}
	expectedBytes := []byte{0x6F, 0x57, 0x3F}
	resultBytes := utils.XorBytesArray(bytes1, bytes2)

	for i, expectedByte := range expectedBytes {
		resultByte := resultBytes[i]
		if resultByte != expectedByte {
			t.Errorf("Bytes doesn't match 0x%x != 0x%x", expectedByte, resultByte)
		}
	}
}
