package main

import (
	"fmt"
	"log"
	"net"
)

func deviceProxy(localAddr, remoteAddr *net.UDPAddr, prefix string, in <-chan []byte, out chan<- []byte) error {
	log.Printf("deviceProxy: %s listen on %v and send to %v", prefix, localAddr, remoteAddr)

	listenConn, err := net.ListenUDP("udp", localAddr)
	if err != nil {
		return fmt.Errorf("could listen on UPD %v", err)
	}
	defer func() {
		err = listenConn.Close()
		if err != nil {
			log.Printf("Error: deviceProxy: fail on close: %v", err)
		}
	}()

	//receive
	go func() {
		buf := make([]byte, 1024)
		for {
			n, _, err := listenConn.ReadFromUDP(buf[0:])
			if err != nil {
				log.Printf("Error: deviceProxy: UDP read error: %v", err)
				continue
			}
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
			_, err := listenConn.WriteToUDP(packet, remoteAddr)
			if err != nil {
				log.Printf("Error: deviceProxy: UDP write error: %v", err)
				continue
			}
		}
	}()

	return nil
}
