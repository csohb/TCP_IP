package main

import (
	"fmt"
	"io"
	"net"
)

func handler(conn net.Conn) {
	recv := make([]byte, 1024)

	for {
		n, err := conn.Read(recv)
		if err != nil {
			if err == io.EOF {
				fmt.Println("connection is closed from client : ", conn.RemoteAddr().String())
			}
			fmt.Println("failed to receive data : ", err)
			break
		}
		if n > 0 {
			fmt.Println("recv : ", recv)
		}
		count := int(recv[0])
		fmt.Println("count : ", count)
		operator := string(recv[count+1 : count+2])
		fmt.Println("operator : ", operator)

		var val int
		switch operator {
		case "+":
			for i := 1; i <= count; i++ {
				val += int(recv[i])
			}
		case "*":
			for i := 1; i <= count; i++ {
				val *= int(recv[i])
			}
		case "-":
			for i := 1; i <= count; i++ {
				val -= int(recv[i])
			}
		}

		fmt.Println("val : ", val)

		send := make([]byte, val)
		send = append(send, byte(val))

		_, err = conn.Write(send)
		if err != nil {
			fmt.Errorf("write send err : %+v", err)
		}
	}
}

func main() {

	listen, err := net.Listen("tcp", ":10002")
	if err != nil {
		fmt.Errorf("failed to listen server : %+v", err)
	}
	defer listen.Close()

	// accept 어떻게 할 것인지 생각.
	/*
		1. 처음에 몇 개의 숫자가 들어 올 것인지.
		2. 받은 숫자 정보 순서대로 저장
		3. 연산자 정보 저장
		4. 연산
		5. 연산결과 client에 전송
	*/

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Failed to Accept : ", err)
			continue
		}
		fmt.Println("accept 완료.")
		go handler(conn)
	}
}
