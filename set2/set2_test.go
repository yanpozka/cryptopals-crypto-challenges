package set2

import (
	"bytes"
	"crypto/aes"
	"encoding/base64"
	"io/ioutil"
	"os"
	"testing"
)

func TestAddPaddingCh9(t *testing.T) {
	const msg = "YELLOW SUBMARINE"

	r := addPadding([]byte(msg), 4)

	expected := []byte("YELLOW SUBMARINE\x04\x04\x04\x04")

	if !bytes.Equal(r, expected) {
		t.Fatalf("%s", r)
	}
}

func TestEncryptDecryptECBModeCh10(t *testing.T) {
	const (
		key  = "YELLOW SUBMARINE"
		iv   = "\x00\x00\x00"
		text = "I Fall in Love Too Easily. I fall in love too fast. I guess."
	)

	cipherAES, err := aes.NewCipher([]byte(key))
	if err != nil {
		t.Fatal(err)
	}

	enc := encryptCBCMode([]byte(text), []byte(iv), cipherAES)

	dec := decryptCBCModeIV(enc, []byte(iv), cipherAES)

	if string(dec[:len(dec)-4]) != text { // -4 coz that's the padding for text
		t.Logf("%q != %q", string(dec), text)
	}

	t.Logf("%q", string(dec))
}

func TestECBModeCh10(t *testing.T) {
	const key = "YELLOW SUBMARINE"
	iv := bytes.Repeat([]byte{byte(0)}, 16)

	f, err := os.Open("10.txt")
	if err != nil {
		t.Fatal(err)
	}

	data, err := ioutil.ReadAll(base64.NewDecoder(base64.StdEncoding, f))
	if err != nil {
		t.Fatal(err)
	}

	cipherAES, err := aes.NewCipher([]byte(key))
	if err != nil {
		t.Fatal(err)
	}

	dec := decryptCBCModeIV(data, []byte(iv), cipherAES)

	t.Logf("%v\n", string(dec))
}
