// Author: [Hilson-Alex](https://github.com/Hilson-Alex).

package exercises

import (
	"io"
)

// ROT13Reader reads from its inner reader and encodes it
// based on the rot13 pattern, as asked in https://go.dev/tour/methods/23.
type ROT13Reader struct {
	reader io.Reader
}

// interface method.
func (reader *ROT13Reader) Read(buffer []byte) (n int, err error) {
	var aux = make([]byte, len(buffer))
	n, err = reader.reader.Read(aux)
	if err == io.EOF {
		return
	}
	for index, elem := range aux {
		aux[index] = rotate(elem)
	}
	copy(buffer, aux)
	return
}

// NewROT13 creates a new ROT13Reader pointer.
func NewROT13(reader io.Reader) *ROT13Reader {
	return &ROT13Reader{
		reader: reader,
	}
}

// isLetter verify if a byte is an ASCII letter. If it is a letter,
// it also returns 'a' for lower case and 'A' for upper case.
func isLetter(b byte) (bool, byte) {
	if b >= byte('a') && b <= byte('z') {
		return true, 'a'
	}
	if b >= byte('A') && b <= byte('Z') {
		return true, 'A'
	}
	return false, 0
}

// rotate a letter by 13 positions
func rotate(b byte) byte {
	var validLetter, reference = isLetter(b)
	var rotated byte = 13
	if !validLetter {
		return b
	}
	rotated += b - reference
	if rotated > 25 {
		rotated -= 26
	}
	return rotated + reference
}
