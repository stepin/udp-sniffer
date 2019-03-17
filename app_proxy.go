package main

import (
	"fmt"
	"log"
	"net"
	"sync"
)

func appProxy(localAddr *net.UDPAddr, prefix string, in <-chan []byte, out chan<- []byte) error {
	log.Printf("appProxy: %s listen on %v and send to unknown yet", prefix, localAddr)
	log.Printf("appProxy: app's IP and port will be autodetected from the first data packet")

	listenConn, err := net.ListenUDP("udp", localAddr)
	if err != nil {
		return fmt.Errorf("could listen on UPD %v, %v", localAddr, err)
	}
	defer func() {
		err = listenConn.Close()
		if err != nil {
			log.Printf("Error: appProxy: fail on close: %v", err)
		}
	}()

	var mutex = &sync.Mutex{}
	var lastClientAddr *net.UDPAddr

	//receive
	go func() {
		buf := make([]byte, 1024)
		for {
			n, address, err := listenConn.ReadFromUDP(buf[0:])
			if err != nil {
				log.Printf("Error: appProxy: UDP read error: %v", err)
				continue
			}

			mutex.Lock()
			if lastClientAddr != address {
				lastClientAddr = address
				log.Printf("Last app address: %v", lastClientAddr)
			}
			mutex.Unlock()

			if n > 0 {
				data := buf[0:n]
				log.Println(prefix + bytes2hex(data))

				out <- data
			}
		}
	}()

	//send
	go func() {
		for {
			packet := <-in

			mutex.Lock()
			address := lastClientAddr
			mutex.Unlock()

			if address == nil {
				log.Printf("Error: appProxy: unknown remote address, packet skipped")
				continue
			}
			_, err := listenConn.WriteToUDP(packet, address)
			if err != nil {
				log.Printf("Error: appProxy: UDP write error: %v", err)
				continue
			}
		}
	}()

	return nil
}
