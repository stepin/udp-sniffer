package main

import (
	"log"
	"net"
)

func udpClientProxyConnection(localAddressString string, remoteAddressString string, prefix string, in <-chan []byte, out chan<- []byte) {
	log.Printf("udpClientProxyConnection: " + prefix + " listen on " + localAddressString + " and send to " + remoteAddressString)

	localAddress, err := net.ResolveUDPAddr("udp", localAddressString)
	checkError(err)

	remoteAddress, err := net.ResolveUDPAddr("udp", remoteAddressString)
	checkError(err)

	listenConn, err := net.ListenUDP("udp", localAddress)
	checkError(err)
	defer listenConn.Close()

	//read goroutine
	go func() {
		buf := make([]byte, 1024)
		for {
			n, _, err := listenConn.ReadFromUDP(buf[0:])
			if err != nil {
				log.Println("Error: UDP read error: ", err)
				continue
			}
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
		_, err := listenConn.WriteToUDP(packet, remoteAddress)
		if err != nil {
			log.Println("Error: UDP write error: ", err)
			continue
		}
	}
}
