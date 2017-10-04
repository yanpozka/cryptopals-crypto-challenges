package set2

import (
	"bytes"
	"crypto/cipher"
)

func addPadding(data []byte, size int) []byte {
	if size < 0 || size > 255 {
		panic("size must be between 0 and 255")
	}

	result := make([]byte, len(data))
	copy(result, data)

	result = append(result, bytes.Repeat([]byte{byte(size)}, size)...)
	return result
}

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
