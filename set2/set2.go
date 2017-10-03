package set2

import "bytes"

func addPadding(data []byte, size int) []byte {
	if size < 0 || size > 255 {
		panic("size must be between 0 and 255")
	}

	pad := bytes.Repeat([]byte{byte(size)}, size)
	result := make([]byte, len(data), len(data)+len(pad))

	copy(result, data)
	result = append(result, pad...)

	return result
}
