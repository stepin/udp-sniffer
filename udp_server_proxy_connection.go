package main

import (
	"log"
	"net"
	"sync"
)

func udpServerProxyConnection(receiveAddressString, prefix string, in <-chan []byte, out chan<- []byte) {
	log.Printf("udpServerProxyConnection: " + prefix + " listen on " + receiveAddressString + " and send to unknown yet")

	localAddress, err := net.ResolveUDPAddr("udp", receiveAddressString)
	checkError(err)

	listenConn, err := net.ListenUDP("udp", localAddress)
	checkError(err)
	defer func() { _ = listenConn.Close() }()

	var mutex = &sync.Mutex{}
	var lastClientAddress *net.UDPAddr

	//read goroutine
	go func() {
		buf := make([]byte, 1024)
		for {
			n, address, err := listenConn.ReadFromUDP(buf[0:])
			if err != nil {
				log.Println("Error: UDP read error: ", err)
				continue
			}

			mutex.Lock()
			if lastClientAddress != address {
				lastClientAddress = address
				log.Println("last client address: ", lastClientAddress)
			}
			mutex.Unlock()

			if n > 0 {
				data := buf[0:n]
				log.Println(prefix + " " + bytesToHexString(data))

				out <- data
			}
		}
	}()

	//write local goroutine
	for {
		packet := <-in

		mutex.Lock()
		address := lastClientAddress
		mutex.Unlock()

		if address == nil {
			log.Print("Error: unknown remote address, packet skipped")
			continue
		}
		_, err := listenConn.WriteToUDP(packet, address)
		if err != nil {
			log.Println("Error: UDP write error: ", err)
			continue
		}
	}
}
