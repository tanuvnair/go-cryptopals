package main

import (
	"cryptopals/sets/set1"
	"fmt"
)

func main() {
	input := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	output := set1.HexToBase64(input)

	fmt.Printf("%s", output)
}
