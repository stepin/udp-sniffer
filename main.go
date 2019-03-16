package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func main() {
	receivePort := flag.Int("receivePort", 11000, "Local UPD port to listen connection")
	receiveIP := flag.String("receiveIP", "", "Local IP port to listen connection")

	localPort := flag.Int("localPort", 11010, "Local UPD port to send data from")
	localIP := flag.String("localIP", "", "Local IP port to send data from")

	remotePort := flag.Int("remotePort", 11000, "Remote UPD port to send data")
	remoteIP := flag.String("remoteIP", "127.0.0.1", "Remote IP port to send data")

	versionFlag := flag.Bool("v", false, "Show app version")
	flag.Parse()

	if *versionFlag {
		printVersion()
		os.Exit(0)
	}

	receiveAddr := *receiveIP + ":" + strconv.Itoa(*receivePort)
	localAddr := *localIP + ":" + strconv.Itoa(*localPort)
	remoteAddr := *remoteIP + ":" + strconv.Itoa(*remotePort)

	log.Println("Started...")
	server(receiveAddr, localAddr, remoteAddr)
	waitTerm()
	log.Println("Stopped...")
}

//Starts required UDP servers and clients.
func server(receiveAddr, localAddr, remoteAddr string) {
	toClient := make(chan []byte, 100)
	fromClient := make(chan []byte, 100)

	err := udpServerProxy(receiveAddr, ">>", fromClient, toClient)
	if err != nil {
		log.Fatalln("Fail on UPD Server start", err)
	}

	err = udpClientProxy(localAddr, remoteAddr, "<<", toClient, fromClient)
	if err != nil {
		log.Fatalln("Fail on UPD Client start", err)
	}
}

//Blocks main func until user press Ctrl+C.
func waitTerm() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	<-signals
}
