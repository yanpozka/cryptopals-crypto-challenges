package set2

import (
	"bytes"
	"testing"
)

func TestAddPaddingCh9(t *testing.T) {
	const msg = "YELLOW SUBMARINE"

	r := addPadding([]byte(msg), 4)

	expected := []byte("YELLOW SUBMARINE\x04\x04\x04\x04")

	if !bytes.Equal(r, expected) {
		t.Fatalf("%+v", r)
	}
}
