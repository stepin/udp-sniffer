package main

import (
	"flag"
)

var localPort int
var localIP string

var receivePort int
var receiveIP string

var remotePort int
var remoteIP string

func parseArgs() {
	flag.IntVar(&receivePort, "receivePort", 11245, "Local UPD port to listen connection")
	flag.StringVar(&receiveIP, "receiveIP", "127.0.0.1", "Local IP port to listen connection")

	flag.IntVar(&localPort, "localPort", 11240, "Local UPD port to send data from")
	flag.StringVar(&localIP, "localIP", "127.0.0.1", "Local IP port to send data from")

	flag.IntVar(&remotePort, "remotePort", 11246, "Remote UPD port to send data")
	flag.StringVar(&remoteIP, "remoteIP", "127.0.0.1", "Remote IP port to send data")

	flag.Parse()
}
