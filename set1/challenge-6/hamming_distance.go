package main

func hammingDistByte(x byte, y byte) int {
	var c int
	a, b := uint8(x), uint8(y)

	for a > 0 || b > 0 {
		if a&1 != b&1 {
			c++
		}
		a >>= 1
		b >>= 1
	}

	return c
}

func hammingDistStr(a, b []byte) int {
	if len(a) > len(b) {
		a, b = b, a
	}
	var dist int
	for ix := range a {
		dist += hammingDistByte(a[ix], b[ix])
	}

	return dist
}
