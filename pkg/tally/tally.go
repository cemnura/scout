package tally

import (
	"bytes"
)

// Finds the count of bytes in given byte
func Tally(bs []byte, sbs []byte) (count int, err error) {
	return TallyCaseSensitive(bs, sbs, true)
}

// Finds the count of bytes in given bytes
func TallyCaseSensitive(bs, sbs []byte, caseSensitiveFlag bool) (count int, err error) {

	if !caseSensitiveFlag {
		bs = bytes.ToLower(bs)
		sbs = bytes.ToLower(sbs)
	}

	count = bytes.Count(bs, sbs)
	return
}