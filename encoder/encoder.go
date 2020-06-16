package encoder

import (
	"encoding/hex"

	"github.com/pankrator/ebml-parser/tools"
)

func Encode(el tools.ElementData) []byte {
	hexData := el.Hex[2:]
	data, err := hex.DecodeString(hexData)
	if err != nil {
		panic(err)
	}
	return data
}

func WriteVInt(value uint64) []byte {
	length := 1
	for length = 1; length <= 8; length += 1 {
		if value < 1<<(7*length)-1 {
			break
		}
	}

	buffer := make([]byte, length)
	val := value
	for i := 1; i <= length; i += 1 {
		b := val & 0xff
		buffer[length-i] = byte(b)
		val -= b
		val /= 1 << 8
	}
	buffer[0] |= 1 << (8 - length)

	return buffer
}
