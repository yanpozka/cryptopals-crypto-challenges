package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("8.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(f)

	for line, err := reader.ReadString('\n'); err == nil; line, err = reader.ReadString('\n') {

		data, decErr := hex.DecodeString(line[:len(line)-1])
		if decErr != nil {
			panic(decErr)
		}

		dupB, isECB := isECBMode(data)
		if isECB {
			fmt.Printf("Duplicated block:\n%s\nIn line:\n%s\n", hex.Dump([]byte(dupB)), hex.Dump(data))
		}
	}
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
