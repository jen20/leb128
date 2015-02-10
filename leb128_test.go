package leb128

import (
	"bytes"
	"testing"
)

var decodingTests = []struct {
	in  []byte
	out uint32
}{
	{[]byte{0x00}, 0},
	{[]byte{0x01}, 1},
	{[]byte{0x05}, 5},
	{[]byte{0xFF, 0xFF, 0xFF, 0xFF, 0x07}, 2147483647},
	{[]byte{0x80, 0x01}, 128},
	{[]byte{0xE5, 0x8E, 0x26}, 624485},
}

func TestDecode(t *testing.T) {
	for _, test := range decodingTests {
		reader := bytes.NewReader(test.in)
		decoded, err := Decode(reader)
		if err != nil {
			t.Fatalf("Error decoding LEB128 to uint32 (Input: %d): %s", test.out, err)
		}
		if decoded != test.out {
			t.Errorf("Expected: %d, Actual: %d", test.out, decoded)
		}
	}
}

var invalidInputTests = [][]byte{
	{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x7}, //Too long
	{}, //EOF
}

func TestDecodeWithInvalidInputs(t *testing.T) {
	for _, test := range invalidInputTests {
		reader := bytes.NewReader(test)
		_, err := Decode(reader)
		if err == nil {
			t.Error("Expected error decoding", test)
		}
	}
}
