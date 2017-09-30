package xorkey

import (
	"sort"
)

func FindSingleXORKey(block []byte) byte {
	maxScore, key := -1, byte(0)

	for k, lim := byte(0), byte(254); k <= lim; k++ {

		pr := make([]byte, len(block))
		for ix := range block {
			pr[ix] = block[ix] ^ k
		}

		if s := score(pr); s > maxScore {
			maxScore = s
			key = k
		}
	}

	return key
}

func score(data []byte) int {
	// e t a o i n s h r d l  u  c  m  f
	// 0 1 2 3 4 5 6 7 8 9 10 11 12 13 14
	//
	// more here: http://letterfrequency.org
	//
	// so far just counting spaces

	// parts := bytes.Split(data, []byte(" "))
	// for _, word := range parts {
	// 	if !isReadable(word) {
	// 		return -1
	// 	}
	// }
	// return len(parts)
	var c int
	for ix := range data {
		if data[ix] == ' ' {
			c++
		}
	}
	return c
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
