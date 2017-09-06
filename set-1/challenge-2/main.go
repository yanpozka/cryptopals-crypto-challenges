//
//  http://cryptopals.com/sets/1/challenges/2
//

package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	const (
		a = "1c0111001f010100061a024b53535009181c"
		b = "686974207468652062756c6c277320657965"
	)

	data, err := xor(a, b)

	fmt.Println(hex.EncodeToString(data), string(data), err)
}

func xor(a, b string) ([]byte, error) {
	if len(a) != len(b) {
		return nil, fmt.Errorf("length of first slice: %d != length of second slice: %d", len(a), len(b))
	}

	dA, err := hex.DecodeString(a)
	if err != nil {
		return nil, err
	}
	dB, err := hex.DecodeString(b)
	if err != nil {
		return nil, err
	}

	result := make([]byte, len(dA))

	for ix := range dA {
		result[ix] = dA[ix] ^ dB[ix]
	}

	return result, nil
}
