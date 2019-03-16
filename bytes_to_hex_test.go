package main

import (
	"testing"
)

func TestByteToHexString(t *testing.T) {
	var tt = []struct {
		name string
		in   []byte
		out  string
	}{
		{"1 byte", []byte{2}, "02 "},
		{"2 bytes", []byte{2, 3}, "02 03 "},
		{"3 bytes", []byte{2, 3, 4}, "02 03 04 "},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			r := bytesToHexString(tc.in)
			if r != tc.out {
				t.Errorf("got %q, want %q", r, tc.out)
			}
		})
	}
}
