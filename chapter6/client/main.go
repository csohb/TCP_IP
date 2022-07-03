package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	remoteAddr, err := net.ResolveUDPAddr("udp", "localhost:6000")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.DialUDP("udp", nil, remoteAddr)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Remote UDP addr : ", conn.RemoteAddr().String())
	log.Println("Local UDP client address : ", conn.LocalAddr().String())

	defer conn.Close()

	msg := []byte("Hello UDP Server!")
	_, err = conn.Write(msg)
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 1024)
	n, addr, err := conn.ReadFromUDP(buf)

	fmt.Println("UDP Server : ", addr)
	fmt.Println("Received from UDP server : ", string(buf[:n]))
}
