package main

import (
	"flag"
	"os"
)

var localPort int
var localIP string

var receivePort int
var receiveIP string

var remotePort int
var remoteIP string

func parseArgs() {
	var versionFlag bool
	flag.BoolVar(&versionFlag, "v", false, "Show app version")

	flag.IntVar(&receivePort, "receivePort", 11000, "Local UPD port to listen connection")
	flag.StringVar(&receiveIP, "receiveIP", "", "Local IP port to listen connection")

	flag.IntVar(&localPort, "localPort", 11010, "Local UPD port to send data from")
	flag.StringVar(&localIP, "localIP", "", "Local IP port to send data from")

	flag.IntVar(&remotePort, "remotePort", 11000, "Remote UPD port to send data")
	flag.StringVar(&remoteIP, "remoteIP", "127.0.0.1", "Remote IP port to send data")

	flag.Parse()

	if versionFlag {
		printVersion()
		os.Exit(0)
	}
}
