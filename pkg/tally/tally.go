package tally

import (
	"bytes"
)

// Finds the count of bytes in given byte
func Tally(bs []byte, sbs []byte) (count int, err error) {
	return TallyCaseSensitive(bs, sbs, false)
}

// Finds the count of bytes in given bytes
func TallyCaseSensitive(bs, sbs []byte, caseInSensitiveFlag bool) (count int, err error) {

	if caseInSensitiveFlag {
		bs = bytes.ToLower(bs)
		sbs = bytes.ToLower(sbs)
	}

	count = bytes.Count(bs, sbs)
	return
}
