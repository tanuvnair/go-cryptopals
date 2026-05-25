package set1

import "math"

var englishFreq = map[byte]float64{
	'a': 0.08167, 'b': 0.01492, 'c': 0.02782, 'd': 0.04253,
	'e': 0.12702, 'f': 0.02228, 'g': 0.02015, 'h': 0.06094,
	'i': 0.06966, 'j': 0.00153, 'k': 0.00772, 'l': 0.04025,
	'm': 0.02406, 'n': 0.06749, 'o': 0.07507, 'p': 0.01929,
	'q': 0.00095, 'r': 0.05987, 's': 0.06327, 't': 0.09056,
	'u': 0.02758, 'v': 0.00978, 'w': 0.02360, 'x': 0.00150,
	'y': 0.01974, 'z': 0.00074,
}

func scoreEnglish(data []byte) float64 {
	var score float64

	for _, b := range data {
		switch {
		case b >= 'A' && b <= 'Z':
			score += englishFreq[b+32]
		case b >= 'a' && b <= 'z':
			score += englishFreq[b]
		case b == ' ':
			score += 0.15
		case b < 32 || b > 126:
			score -= 5
		}
	}

	return score
}

func BruteforceXorDecipher(cipherText []byte) []byte {
	bestScore := math.Inf(-1)
	var bestPlaintext []byte

	for key := 0; key < 256; key++ {
		candidate := make([]byte, len(cipherText))
		for i := range cipherText {
			candidate[i] = cipherText[i] ^ byte(key)
		}

		score := scoreEnglish(candidate)
		if score > bestScore {
			bestScore = score
			bestPlaintext = candidate
		}
	}

	return bestPlaintext
}
