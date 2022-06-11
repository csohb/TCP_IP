package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	// dial & connect to server
	conn, err := net.Dial("tcp", ":10000")
	if err != nil {
		fmt.Println("Failed to Dial : ", err)
	}
	defer conn.Close()

	// go routine 익명함수
	go func(c net.Conn) {
		// 보낼 메시지
		send := "Hello"
		for {
			// connection에 보낼 메시지 write
			_, err = c.Write([]byte(send))
			if err != nil {
				fmt.Println("Failed to write data : ", err)
				break
			}
			// 1초마다 반복
			time.Sleep(time.Second * 1)
		}
		/*data := make([]byte, 4096)

		for {
			n, err := conn.Read(data)
			if err != nil {
				fmt.Println("Client Read err : ", err)
				return
			}

			fmt.Println("Server send : ", string(data[:n]))
			time.Sleep(time.Duration(3)*time.Second)
		}*/
	}(conn)

	go func(c net.Conn) {
		// connection으로 부터 전달 받은 msg 저장할 buffer
		recv := make([]byte, 4096)

		for {
			// read msg from conn
			n, err := c.Read(recv)
			if err != nil {
				fmt.Println("Failed to Read data : ", err)
				break
			}
			fmt.Println("Recv data : ", string(recv[:n]))
		}
	}(conn)

	// 사용자 입력이 들어올때까지 blocking 했다가 입력을 마치면 서버로 전송
	for {
		var s string
		fmt.Scanln(&s)
		conn.Write([]byte(s))
		time.Sleep(time.Duration(3) * time.Second)
	}
}
