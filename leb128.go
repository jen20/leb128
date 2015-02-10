package leb128

import (
	"encoding/binary"
	"errors"
	"io"
)

const (
	errMaximumEncodingLengthExceeded = "Exceeded maximum of 5 bytes for a LEB128-encoded uint32."
)

//Decode reads a LEB128-encoded unsigned integer from the given io.Reader and returns it
//as a uint32.
func Decode(reader io.Reader) (uint32, error) {
	var result uint32
	var shift uint8
	var current byte
	for {
		if err := binary.Read(reader, binary.LittleEndian, &current); err != nil {
			return 0, err
		}
		result |= uint32(current&0x7F) << shift
		if (current & 0x80) == 0 {
			break
		}
		shift += 7
		if shift > (4 * 7) {
			return 0, errors.New(errMaximumEncodingLengthExceeded)
		}
	}
	return result, nil
}
