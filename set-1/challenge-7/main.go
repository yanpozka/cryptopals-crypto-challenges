package main

import (
	"crypto/aes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	f, err := os.Open("7.txt")
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadAll(base64.NewDecoder(base64.StdEncoding, f))
	if err != nil {
		panic(err)
	}

	cipher, err := aes.NewCipher([]byte("YELLOW SUBMARINE"))
	if err != nil {
		panic(err)
	}

	block := make([]byte, cipher.BlockSize())

	for ix := 0; ix < len(data); ix += cipher.BlockSize() {
		cipher.Decrypt(block, data[ix:ix+cipher.BlockSize()])

		fmt.Printf("%s", block)
	}
}
