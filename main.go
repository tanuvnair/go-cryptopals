package main

import (
	"cryptopals/sets/set1"
	"fmt"
)

func main() {
	input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

	input1DecodedHex, _ := set1.DecodeHex([]byte(input))
	decipheredPlainText := set1.BruteforceXorDecipher(input1DecodedHex)

	fmt.Printf("%s \n", decipheredPlainText)
}
