package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// simple function to verify error
func checkError(err error) {
	if err != nil {
		log.Fatalln("Error: ", err)
	}
}

//blocks main func until Ctrl+C
func waitingForTerminalSignals() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	<-signals
}

//convert bytes array to a hex string
func bytesToHexString(data []byte) string {
	msg := ""
	for _, aByte := range data {
		msg += fmt.Sprintf("%02x ", aByte)
	}
	return msg
}
