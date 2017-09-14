package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
)

func main() {
	fmt.Printf("distance %d == 37\n", hammingDistStr([]byte("this is a test"), []byte("wokka wokka!!!")))

	f, err := os.Open("data-6.txt")
	if err != nil {
		panic(err)
	}

	dr := base64.NewDecoder(base64.StdEncoding, bufio.NewReader(f))
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
	var norm []pair

	for keySize := 2; keySize <= 40; keySize++ {
		a, b := data[:keySize], data[keySize+1:(keySize*2)%len(data)]

		norm = append(norm, pair{size: keySize, val: float64(hammingDistStr(a, b)) / float64(keySize)})
	}

	sort.Slice(norm[:], func(i, j int) bool { return norm[i].val < norm[j].val })

	fmt.Println(norm)
}
