package main

import (
	"cryptopals/sets/set1"
	"fmt"
)

func main() {
	input1 := "1c0111001f010100061a024b53535009181c"
	input2 := "686974207468652062756c6c277320657965"

	input1DecodedHex, _ := set1.DecodeHex([]byte(input1))
	input2DecodedHex, _ := set1.DecodeHex([]byte(input2))
	output := set1.Xor(input1DecodedHex, input2DecodedHex)

	fmt.Printf("Encoded Hex Output: %s \n", set1.EncodeHex(output))
	fmt.Printf("Formatted Output: %s \n", string(output))
}
