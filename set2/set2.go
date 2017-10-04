package set2

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	mathRand "math/rand"
	"time"
)

// challenge 9

func addPadding(data []byte, size int) []byte {
	if size < 0 || size > 255 {
		panic("size must be between 0 and 255")
	}

	result := make([]byte, len(data))
	copy(result, data)

	result = append(result, bytes.Repeat([]byte{byte(size)}, size)...)
	return result
}

// challenge 10

func xor(key, block []byte) []byte {
	r := make([]byte, len(block))
	for ix := range block {
		r[ix] = block[ix] ^ key[ix%len(key)]
	}
	return r
}

func encryptCBCMode(data, initVector []byte, cb cipher.Block) []byte {
	if mod := len(data) % cb.BlockSize(); mod != 0 {
		data = addPadding(data, cb.BlockSize()-mod)
	}

	buff := new(bytes.Buffer)
	section := make([]byte, cb.BlockSize())
	prev := initVector

	for ix := 0; ix < len(data); ix += cb.BlockSize() {
		currBlock := data[ix : ix+cb.BlockSize()]

		// XOR with previous block before encrypt
		xorb := xor(prev, currBlock)

		cb.Encrypt(section, xorb)

		buff.Write(section)
		prev = section
	}

	return buff.Bytes()
}

func decryptCBCModeIV(data, initVector []byte, cb cipher.Block) []byte {
	buff := new(bytes.Buffer)
	prev := initVector
	section := make([]byte, cb.BlockSize())

	for ix := 0; ix < len(data); ix += cb.BlockSize() {
		currBlock := data[ix : ix+cb.BlockSize()]

		// inverse order, first decrypt
		cb.Decrypt(section, currBlock)

		// xor with previous block
		buff.Write(xor(prev, section))
		prev = currBlock
	}

	return buff.Bytes()
}

// challenge 11

func encryptECBMode(data []byte, cb cipher.Block) []byte {
	if mod := len(data) % cb.BlockSize(); mod != 0 {
		data = addPadding(data, cb.BlockSize()-mod)
	}

	block := make([]byte, cb.BlockSize())
	buff := new(bytes.Buffer)

	for ix := 0; ix < len(data); ix += cb.BlockSize() {
		cb.Encrypt(block, data[ix:ix+cb.BlockSize()])
		buff.Write(block)
	}

	return buff.Bytes()
}

func encryptionOracle(in []byte) []byte {
	mathRand.Seed(time.Now().UnixNano())

	key := make([]byte, aes.BlockSize)
	rand.Read(key)

	extra := make([]byte, mathRand.Intn(5)+5)
	rand.Read(extra)
	println("extra prefix and suffix len: ", len(extra))

	buff := new(bytes.Buffer)
	buff.Write(extra)
	buff.Write(in)
	buff.Write(extra)

	data := buff.Bytes()

	cipherAES, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// ECB mode
	if mathRand.Intn(2) == 0 {
		println("ECB mode")
		return encryptECBMode(data, cipherAES)
	}

	// CBC mode
	iv := make([]byte, aes.BlockSize)
	rand.Read(iv)
	println("CBC mode")
	return encryptCBCMode(data, iv, cipherAES)
}
