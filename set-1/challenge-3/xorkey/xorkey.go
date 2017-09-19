package xorkey

import (
	"fmt"
	"sort"
	"unicode"
)

func FindSingleXORKey(block []byte) byte {
	maxScore, key := -1, byte(0)

	for k, lim := byte('A'), byte('Z'); k <= lim; k++ {

		pr := make([]byte, len(block))
		for ix := range block {
			pr[ix] = block[ix] ^ k
		}
		fmt.Println(pr)

		if isReadable(string(pr)) {
			if s := score(pr); s > maxScore {
				maxScore = s
				key = k
			}
		}
	}

	return key
}

func score(data []byte) (counter int) {
	// e t a o i n s h r d l  u  c  m  f
	// 0 1 2 3 4 5 6 7 8 9 10 11 12 13 14
	//
	// more here: http://letterfrequency.org
	//
	// so far just counting letters and spaces

	for _, b := range data {
		if b == ' ' {
			counter += 3
			continue
		}

		if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') {
			counter++
		}
	}

	return counter
}
func isReadable(data string) bool {
	for _, d := range data {
		if !unicode.IsPrint(d) {
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
