package main

import (
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func main() {
	parseArgs()

	if versionFlag {
		printVersion()
		os.Exit(0)
	}

	log.Println("Started...")

	receiveAddress := receiveIP + ":" + strconv.Itoa(receivePort)
	localAddress := localIP + ":" + strconv.Itoa(localPort)
	remoteAddress := remoteIP + ":" + strconv.Itoa(remotePort)

	toClient := make(chan []byte, 100)
	fromClient := make(chan []byte, 100)

	err := udpServerProxyConnection(receiveAddress, ">>", fromClient, toClient)
	if err != nil {
		log.Fatalln("Fail on UPD Server start", err)
	}
	err = udpClientProxyConnection(localAddress, remoteAddress, "<<", toClient, fromClient)
	if err != nil {
		log.Fatalln("Fail on UPD Client start", err)
	}

	waitingForTerminalSignals()
	log.Println("Stopped...")
}

//Blocks main func until user press Ctrl+C.
func waitingForTerminalSignals() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	<-signals
}
