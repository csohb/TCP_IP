package main

import (
	"fmt"
	"io"
	"net"
)

func handler(conn net.Conn) {
	// 받은 msg 저장할 buffer
	recv := make([]byte, 4096)

	for {
		// read msg from client
		n, err := conn.Read(recv)
		if err != nil {
			if err == io.EOF {
				fmt.Println("connection is closed from client : ", conn.RemoteAddr().String())
			}
			fmt.Println("Failed to receive data : ", err)
			break
		}
		// msg print
		if n > 0 {
			fmt.Println(string(recv[:n]))
			// connection에 보낼 메시지 write
			//conn.Write(recv[:n])
		}
	}
}

func main() {
	// socket 생성
	listener, err := net.Listen("tcp", ":10000")
	if err != nil {
		fmt.Println("Failed to Listen : ", err)
	}
	// 프로세스가 멈추면 listener도 close
	defer listener.Close()

	// accept
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Failed to Accept : ", err)
			continue
		}

		// accept 한 후 client에서 받은 메시지 handle
		go handler(conn)
	}
}
