package main

import (
	"fmt"
)

//Convert bytes array to a hex string. Space before each byte.
func bytes2hex(data []byte) string {
	s := ""
	for _, b := range data {
		s += fmt.Sprintf(" %02x", b)
	}
	return s
}
