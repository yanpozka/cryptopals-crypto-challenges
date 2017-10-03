package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"sort"

	"github.com/yanpozka/cryptopals-crypto-challenges/set-1/challenge-3/xorkey"
)

func main() {
	// tryKeySize(3, []byte("abcabcab"))
	fmt.Printf("distance %d == 37\n", hammingDistStr([]byte("this is a test"), []byte("wokka wokka!!!")))

	f, err := os.Open("data-6.txt")
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadAll(base64.NewDecoder(base64.StdEncoding, bufio.NewReader(f)))
	if err != nil {
		panic(err)
	}

	//
	repoxr(data)
}

type pair struct {
	val  float64
	size int
}

func repoxr(data []byte) {
	norm := make([]pair, 0, 40)

	for keySize := 2; keySize <= 40; keySize++ {
		a, b := data[:keySize*4], data[keySize*4:(keySize*8)%len(data)]

		norm = append(norm, pair{size: keySize, val: float64(hammingDistStr(a, b)) / float64(keySize*4)})
	}

	sort.Slice(norm, func(i, j int) bool { return norm[i].val < norm[j].val })

	// for ix := 0; ix < 3; ix++ {
	key := tryKeySize(norm[0].size, data)
	// }

	for ix := range data {
		data[ix] ^= key[ix%len(key)]
	}

	fmt.Println("\n", string(data))
}

func tryKeySize(keySize int, data []byte) (key []byte) {

	for k := 0; k < keySize; k++ {
		var block []byte
		for ix := k; ix < len(data); ix += keySize {
			block = append(block, data[ix])
		}

		k, _, _ := xorkey.FindSingleXORKey(block)
		key = append(key, k)
	}

	fmt.Printf("key size= %d key= %q \n", keySize, key)
	return
}
