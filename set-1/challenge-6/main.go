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

	dr := base64.NewDecoder(base64.StdEncoding, bufio.NewReader(f))
	// _, err = ioutil.ReadAll(dr)
	data, err := ioutil.ReadAll(dr)
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
		a, b := data[:keySize], data[keySize+1:(keySize*2)%len(data)]
		norm = append(norm, pair{size: keySize, val: float64(hammingDistStr(a, b)) / float64(keySize)})
	}

	sort.Slice(norm, func(i, j int) bool { return norm[i].val < norm[j].val })

	for ix := 0; ix < 4; ix++ {
		tryKeySize(norm[ix].size, data)
	}
}

func tryKeySize(keySize int, data []byte) {
	var key []byte

	for k := 0; k < keySize; k++ {
		var block []byte
		for ix := k; ix < len(data); ix += keySize {
			block = append(block, data[ix])
		}
		key = append(key, xorkey.FindSingleXORKey(block))
	}

	fmt.Println(keySize, key, len(key))
}
