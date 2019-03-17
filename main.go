package main

import (
	"flag"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func main() {
	proxyAppPort := flag.Int("proxyAppPort", 11000, "UPD port for this tool to communicate with app")
	proxyAppIP := flag.String("proxyAppIP", "", "IP for this tool to communicate with device app")

	proxyDevicePort := flag.Int("proxyDevicePort", 11010, "UPD port for this tool to communicate with device")
	proxyDeviceIP := flag.String("proxyDeviceIP", "", "IP for this tool to communicate with device")

	devicePort := flag.Int("devicePort", 11000, "UPD port of device")
	deviceIP := flag.String("deviceIP", "127.0.0.1", "IP of device")

	versionFlag := flag.Bool("v", false, "Show app version")
	flag.Parse()

	if *versionFlag {
		printVersion()
		os.Exit(0)
	}

	proxyAppAddr := resolveUDPAddr(*proxyAppIP, *proxyAppPort)
	proxyDeviceAddr := resolveUDPAddr(*proxyDeviceIP, *proxyDevicePort)
	deviceAddr := resolveUDPAddr(*deviceIP, *devicePort)

	log.Println("Started...")
	startProxies(proxyAppAddr, proxyDeviceAddr, deviceAddr)
	waitTerm()
	log.Println("Stopped...")
}

func resolveUDPAddr(ip string, port int) *net.UDPAddr {
	s := ip + ":" + strconv.Itoa(port)
	addr, err := net.ResolveUDPAddr("udp", s)
	if err != nil {
		log.Fatalf("Error: could not resolve UPD addr %v, error: %v\n", s, err)
	}
	return addr
}

//Starts required UDP servers and clients.
func startProxies(proxyAppAddr, proxyDeviceAddr, deviceAddr *net.UDPAddr) {
	toDevice := make(chan []byte, 100)
	fromDevice := make(chan []byte, 100)

	//appAddr is autodetected from the first data packet
	err := appProxy(proxyAppAddr, ">>", fromDevice, toDevice)
	if err != nil {
		log.Fatalln("Fail on UPD Server start", err)
	}

	err = deviceProxy(proxyDeviceAddr, deviceAddr, "<<", toDevice, fromDevice)
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
