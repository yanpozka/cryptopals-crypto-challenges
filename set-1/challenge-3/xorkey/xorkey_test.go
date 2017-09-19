package xorkey

import (
	"encoding/hex"
	"testing"
)

func TestFindSingleXORKey(t *testing.T) {
	const s = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	data, err := hex.DecodeString(s)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%q", string(FindSingleXORKey(data)))
}
