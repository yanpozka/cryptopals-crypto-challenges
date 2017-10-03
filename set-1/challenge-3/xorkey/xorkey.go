//
// Idea taken from  https://github.com/FiloSottile/mostly-harmless/blob/master/cryptopals/set1.go
//
package xorkey

import (
	"bufio"
	"io"
	"os"
	"sort"
)

var freqs = loadData()

func loadData() map[byte]float64 {
	f1, err := os.Open("834-0.txt")
	if err != nil {
		panic(err)
	}
	f2, err := os.Open("pg128.txt")
	if err != nil {
		panic(err)
	}

	freq := map[byte]float64{}
	reader := bufio.NewReader(io.MultiReader(f1, f2))

	for b, err := reader.ReadByte(); err == nil; b, err = reader.ReadByte() {
		freq[b]++
	}

	return freq
}

func FindSingleXORKey(block []byte) (key byte, dec []byte) {
	var maxScore float64

	for k, lim := 0, 255; k <= lim; k++ {

		pr := make([]byte, len(block))
		for ix := range block {
			pr[ix] = block[ix] ^ byte(k)
		}

		if s := score(pr); s > maxScore {
			maxScore = s
			key = byte(k)
			dec = pr
		}
	}

	return
}

func score(data []byte) float64 {
	// e t a o i n s h r d l  u  c  m  f
	// 0 1 2 3 4 5 6 7 8 9 10 11 12 13 14
	//
	// more here: http://letterfrequency.org

	var s float64
	for _, b := range data {
		s += freqs[b]
	}
	return s / float64(len(freqs))
}

func isReadable(data []byte) bool {
	for _, b := range data {
		if b < ' ' || b > '~' {
			return false
		}
	}
	return true
}

type Pair struct {
	Freq int
	Val  byte
}

func GetSortedFreqs(data []byte) []Pair {
	freqs := map[byte]int{}

	for _, b := range data {
		freqs[b]++
	}

	orderedFreqs := make([]Pair, 0, len(freqs))
	for b, f := range freqs {
		orderedFreqs = append(orderedFreqs, Pair{Val: b, Freq: f})
	}

	sort.Slice(orderedFreqs, func(i, j int) bool { return orderedFreqs[i].Freq > orderedFreqs[j].Freq })

	return orderedFreqs
}
