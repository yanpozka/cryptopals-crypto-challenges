package set2

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	mathRand "math/rand"
	"net/url"
	"time"
)

// challenge 9

func addPadding(data []byte, size int) []byte {
	if size < 0 || size > 255 {
		panic("size must be between 0 and 255")
	}

	result := make([]byte, len(data))
	copy(result, data)

	return append(result, bytes.Repeat([]byte{byte(size)}, size)...)
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

func encryptionOracle(in []byte, random bool) []byte {
	mathRand.Seed(time.Now().UnixNano())

	key := make([]byte, aes.BlockSize)
	rand.Read(key)

	return _encryptionOracle(in, key, random)
}

func _encryptionOracle(in, key []byte, random bool) []byte {
	extra := make([]byte, mathRand.Intn(5)+5)
	rand.Read(extra)

	buff := new(bytes.Buffer)
	buff.Write(extra)
	buff.Write(in)
	buff.Write(extra)

	data := buff.Bytes()

	cipherAES, err := aes.NewCipher(key)
	panicIfErr(err)

	// ECB mode
	if random == false || mathRand.Intn(2) == 0 {
		// println("ECB mode")
		return encryptECBMode(data, cipherAES)
	}

	// CBC mode
	iv := make([]byte, aes.BlockSize)
	rand.Read(iv)
	// println("CBC mode")
	return encryptCBCMode(data, iv, cipherAES)
}

func isECBMode(data []byte) ([]byte, bool) {
	blocks := make(map[string]struct{})

	for ix := 0; ix < len(data); ix += 16 {
		end := ix + 16
		if end > len(data) {
			end = len(data)
		}

		cb := string(data[ix:end])
		if _, contains := blocks[cb]; contains {
			return []byte(cb), true
		}

		blocks[cb] = struct{}{}
	}

	return nil, false
}

// challenge 12

//
// AES-128-ECB(your-string || unknown-string, random-key)
//
func encrypOracleWithConstKey(in []byte) []byte {
	const (
		key   = "Summer Time live"
		extra = `Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkg
aGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBq
dXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUg
YnkK`
	)

	decB64, err := base64.StdEncoding.DecodeString(extra)
	panicIfErr(err)

	return _encryptionOracle(append(in, decB64...), []byte(key), false)
}

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

// decrypts "unknown-string" from function above
func decryptUnknownStrECB(in []byte) []byte {
	// TODO:
	// 1. Feed identical bytes of your-string to the function, 1 at a time - start with 1 byte "A", then "AA", then "AAA" and so on.
	//    Discover the block size of the cipher.
	//
	const blockSize = aes.BlockSize

	// 2. Detect that the function is using ECB.
	if _, isECB := isECBMode(in); !isECB {
		panic("not ECB mode. WTF!")
	}

	// 3. Knowing the block size, craft an input block that is exactly 1 byte short.
	input := bytes.Repeat([]byte("A"), blockSize)
	println(input)
	// TODO:
	// 4. Make a dictionary of every possible last byte by feeding different strings to the oracle,
	//    remembering the first block of each invocation

	// TODO:
	// 5. Match the output of the one-byte-short input to one of the entries in your dictionary. You've now discovered the first byte of unknown-string ??

	// TODO:
	// 6. Repeat for the next byte.

	return nil
}

// challenge 13

var userIDCounter int

func profileFor(email string) string {
	userIDCounter++

	v := url.Values{}
	v.Set("email", email)
	v.Set("uid", fmt.Sprintf("%d", userIDCounter))
	v.Set("role", "user")

	return v.Encode()
}

func decodeProfile(encoded string) url.Values {
	v, err := url.ParseQuery(encoded)
	panicIfErr(err)
	return v
}

func encryptProfile(encoded string) (key, cipherTxt []byte) {
	key = make([]byte, aes.BlockSize)
	rand.Read(key)

	cipherAES, err := aes.NewCipher(key)
	panicIfErr(err)

	cipherTxt = encryptECBMode([]byte(encoded), cipherAES)
	return
}

// copied from challenge 7
func decryptAESECBMode(key, data []byte) []byte {
	cipher, err := aes.NewCipher(key)
	panicIfErr(err)

	dst := make([]byte, aes.BlockSize)
	buff := new(bytes.Buffer)

	for ix := 0; ix < len(data); ix += aes.BlockSize {
		cipher.Decrypt(dst, data[ix:ix+aes.BlockSize])

		_, err := buff.Write(dst)
		panicIfErr(err)
	}

	return buff.Bytes()
}
