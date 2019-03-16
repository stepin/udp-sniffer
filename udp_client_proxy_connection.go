package main

import (
	"fmt"
	"log"
	"net"
)

func udpClientProxy(localAddrStr, remoteAddrStr, prefix string, in <-chan []byte, out chan<- []byte) error {
	log.Printf("udpClientProxy: %s listen on %s and send to %s", prefix, localAddrStr, remoteAddrStr)

	localAddr, err := net.ResolveUDPAddr("udp", localAddrStr)
	if err != nil {
		return fmt.Errorf("could not resolve UPD addr for local address %v", err)
	}

	remoteAddr, err := net.ResolveUDPAddr("udp", remoteAddrStr)
	if err != nil {
		return fmt.Errorf("could not resolve UPD addr for remote address %v", err)
	}

	listenConn, err := net.ListenUDP("udp", localAddr)
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
				log.Println(prefix + bytes2hex(data))

				out <- data
			}
		}
	}()

	//write local goroutine
	go func() {
		for {
			packet := <-in
			_, err := listenConn.WriteToUDP(packet, remoteAddr)
			if err != nil {
				log.Println("Error: client: UDP write error: ", err)
				continue
			}
		}
	}()

	return nil
}
