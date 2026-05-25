package set1

import "encoding/hex"

func EncodeHex(input []byte) []byte {
	encodedHex := make([]byte, hex.EncodedLen(len(input)))
	hex.Encode(encodedHex, input)

	return encodedHex
}

func Xor(firstInput []byte, secondInput []byte) []byte {
	destination := make([]byte, len(firstInput))

	for i := 0; i < len(firstInput); i++ {
		destination[i] = firstInput[i] ^ secondInput[i]
	}

	return destination
}
