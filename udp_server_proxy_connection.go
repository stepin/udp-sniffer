package main

import (
	"fmt"
	"log"
	"net"
	"sync"
)

func udpServerProxy(receiveAddrStr, prefix string, in <-chan []byte, out chan<- []byte) error {
	log.Printf("udpServerProxy: " + prefix + " listen on " + receiveAddrStr + " and send to unknown yet")

	localAddr, err := net.ResolveUDPAddr("udp", receiveAddrStr)
	if err != nil {
		return fmt.Errorf("could not resolve UPD addr for receiver address %v", err)
	}

	listenConn, err := net.ListenUDP("udp", localAddr)
	if err != nil {
		return fmt.Errorf("could listen on UPD %v, %v", localAddr, err)
	}
	defer func() {
		err = listenConn.Close()
		if err != nil {
			log.Printf("Error: server: fail on close: %v\n", err)
		}
	}()

	var mutex = &sync.Mutex{}
	var lastClientAddr *net.UDPAddr

	//read goroutine
	go func() {
		buf := make([]byte, 1024)
		for {
			n, address, err := listenConn.ReadFromUDP(buf[0:])
			if err != nil {
				log.Printf("Error: server: UDP read error: %v\n", err)
				continue
			}

			mutex.Lock()
			if lastClientAddr != address {
				lastClientAddr = address
				log.Printf("Last client address: %v\n", lastClientAddr)
			}
			mutex.Unlock()

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

			mutex.Lock()
			address := lastClientAddr
			mutex.Unlock()

			if address == nil {
				log.Printf("Error: server: unknown remote address, packet skipped\n")
				continue
			}
			_, err := listenConn.WriteToUDP(packet, address)
			if err != nil {
				log.Printf("Error: server: UDP write error: %v\n", err)
				continue
			}
		}
	}()

	return nil
}
