package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":10001")
	if err != nil {
		fmt.Errorf("client dial failed : %+v", err)
	}
	defer conn.Close()
	var writeBuffer []byte
	var num int
	fmt.Print("how many numbers do you want to send to server? :")
	fmt.Scanln(&num)
	writeBuffer = append(writeBuffer, byte(num))

	for i := 0; i < num; i++ {
		var input int
		fmt.Print("input number you want to operate : ")
		fmt.Scanln(&input)
		writeBuffer = append(writeBuffer, byte(input))
	}

	var operator string
	fmt.Print("input operator : ")
	fmt.Scanln(&operator)
	operatorByte := []byte(operator)
	for i := 0; i < len(operatorByte); i++ {
		writeBuffer = append(writeBuffer, operatorByte[i])
	}
	fmt.Println("writeBuffer : ", writeBuffer)
	// 보낼 숫자 갯수, 연속 숫자, operator 한번에 모아서 전송? write()

	_, err = conn.Write(writeBuffer)
	if err != nil {
		fmt.Errorf("write to Server failed. : %+v", err)
	}

	// 연산결과 read()로 읽어오기

	// connection으로 부터 전달 받은 msg 저장할 buffer
	recv := make([]byte, 4096)
	for {
		// read msg from conn
		n, err := conn.Read(recv)
		if err != nil {
			fmt.Println("Failed to Read data : ", err)
			break
		}
		if n > 0 {
			fmt.Println("Recv data : ", recv[:n])
		}
	}
}
