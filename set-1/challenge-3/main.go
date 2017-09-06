//
//  http://cryptopals.com/sets/1/challenges/3
//
//  answer is 'X'

package main

import (
	"encoding/hex"
	"fmt"
	"unicode"
)

func main() {
	const s = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

	data, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}

	r := make([]byte, len(data))

	for k, lim := byte('A'), byte('Z'); k <= lim; k++ {

		for ix := range data {
			r[ix] = data[ix] ^ k
		}

		// printing the current buffer of bytes may break the output
		//
		if isReadable(string(r)) {
			fmt.Println(string(k), string(r))
		} else {
			fmt.Println(string(k))
		}
	}
}

func isReadable(data string) bool {
	for _, d := range data {
		if !unicode.IsPrint(d) {
			return false
		}
	}
	return true
}

