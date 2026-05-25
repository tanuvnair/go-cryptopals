package set1

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func DecodeHex(input []byte) ([]byte, error) {
	// We allocate the decoded length of the input in decodedHex, we use hex.DecodedLen as the input is in bytes and after decoding, it will take half the input length, i.e two hex characters is one byte
	decodedHexBytes := make([]byte, hex.DecodedLen(len(input)))

	// We write to the allocated space in decodedHex
	_, err := hex.Decode(decodedHexBytes, input)

	if err != nil {
		return nil, err
	}

	return decodedHexBytes, nil
}

func Base64Encode(input []byte) []byte {
	// We allocate space for the encoded length of the input (basically the length after converting the bytes to base64)
	encodedBase64 := make([]byte, base64.StdEncoding.EncodedLen(len(input)))

	// We write to the allocated space
	base64.StdEncoding.Encode(encodedBase64, input)

	return encodedBase64
}

func HexToBase64(hexValue string) []byte {
	// Here we get the hexadecimal string as their numeric byte values, so each character is a number
	hexBytes := []byte(hexValue)

	decodedHexBytes, err := DecodeHex(hexBytes)
	if err != nil {
		fmt.Printf("failed to decode hex: %s", err)
		return nil
	}

	base64Encoded := Base64Encode(decodedHexBytes)

	return base64Encoded
}
