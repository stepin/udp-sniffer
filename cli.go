package main

import (
	"log"
	"strconv"
)

func main() {
	parseArgs()
	log.Println("Started...")

	receiveAddress := receiveIP + ":" + strconv.Itoa(receivePort)
	localAddress := localIP + ":" + strconv.Itoa(localPort)
	remoteAddress := remoteIP + ":" + strconv.Itoa(remotePort)

	toClient := make(chan []byte, 100)
	fromClient := make(chan []byte, 100)

	go udpServerProxyConnection(receiveAddress, ">>", fromClient, toClient)
	go udpClientProxyConnection(localAddress, remoteAddress, "<<", toClient, fromClient)

	waitingForTerminalSignals()
	log.Println("Stopped...")
}
