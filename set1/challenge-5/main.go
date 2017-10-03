//
// http://cryptopals.com/sets/1/challenges/5
//

package main

import (
	"encoding/hex"
	"fmt"
)

const (
	key = "ICE"
	src = `Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal`
)

func main() {

	out := hex.EncodeToString(cipherXOR([]byte(key), []byte(src)))
	fmt.Println(out)

	fmt.Println(out == "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f")
}

func cipherXOR(key, src []byte) []byte {

	r := make([]byte, len(src))

	for ix := range src {
		r[ix] = src[ix] ^ key[ix%len(key)]
	}

	return r
}

