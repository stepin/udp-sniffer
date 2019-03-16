package main

import (
	"fmt"
)

//Convert bytes array to a hex string. Space after each byte.
func bytesToHexString(data []byte) string {
	s := ""
	for _, b := range data {
		s += fmt.Sprintf("%02x ", b)
	}
	return s
}
