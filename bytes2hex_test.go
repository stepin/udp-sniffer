package main

import (
	"testing"
)

func TestByte2hex(t *testing.T) {
	var cases = []struct {
		name string
		in   []byte
		out  string
	}{
		{"1 byte", []byte{2}, " 02"},
		{"2 bytes", []byte{2, 3}, " 02 03"},
		{"3 bytes", []byte{2, 3, 4}, " 02 03 04"},
	}
	t.Parallel()
	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			r := bytes2hex(tc.in)
			if r != tc.out {
				t.Errorf("got %q, want %q", r, tc.out)
			}
		})
	}
}
