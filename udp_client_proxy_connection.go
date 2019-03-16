package main

import (
	"fmt"
	"log"
	"net"
)

func udpClientProxyConnection(localAddressString string, remoteAddressString string, prefix string, in <-chan []byte, out chan<- []byte) error {
	log.Printf("udpClientProxyConnection: " + prefix + " listen on " + localAddressString + " and send to " + remoteAddressString)

	localAddress, err := net.ResolveUDPAddr("udp", localAddressString)
	if err != nil {
		return fmt.Errorf("could not resolve UPD addr for local address %v", err)
	}

	remoteAddress, err := net.ResolveUDPAddr("udp", remoteAddressString)
	if err != nil {
		return fmt.Errorf("could not resolve UPD addr for remote address %v", err)
	}

	listenConn, err := net.ListenUDP("udp", localAddress)
	if err != nil {
		return fmt.Errorf("could listen on UPD %v", err)
	}
	defer func() { _ = listenConn.Close() }()

	//read goroutine
	go func() {
		buf := make([]byte, 1024)
		for {
			n, _, err := listenConn.ReadFromUDP(buf[0:])
			if err != nil {
				log.Println("Error: client: UDP read error: ", err)
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
	go func() {
		for {
			packet := <-in
			_, err := listenConn.WriteToUDP(packet, remoteAddress)
			if err != nil {
				log.Println("Error: client: UDP write error: ", err)
				continue
			}
		}
	}()

	return nil
}
