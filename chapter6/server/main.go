package main

import (
	"fmt"
	"log"
	"net"
)

func handleUDPConnection(conn *net.UDPConn) {
	buffer := make([]byte, 1024)

	n, addr, err := conn.ReadFromUDP(buffer)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("UDP client : ", addr)
	fmt.Println("Received from UDP client : ", string(buffer[:n]))

	msg := []byte("Hello UDP Client")
	_, err = conn.WriteToUDP(msg, addr)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	udpAddr, err := net.ResolveUDPAddr("udp4", "localhost:6000")
	if err != nil {
		log.Fatal(err)
	}

	listen, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("UDP server up and listening on port 6000")
	defer listen.Close()

	for {
		handleUDPConnection(listen)
	}
}


