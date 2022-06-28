package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":10002")
	if err != nil {
		fmt.Errorf("client dial failed : %+v", err)
	}
	defer conn.Close()

	// 보낼 숫자 갯수, 연속 숫자, operator 한번에 모아서 전송? write()
	go func(conn net.Conn) {
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
		_, err = conn.Write(writeBuffer)
		if err != nil {
			fmt.Errorf("write to Server failed. : %+v", err)
		}
	}(conn)

	// 연산결과 read()로 읽어오기
	go func(conn net.Conn) {
		readByte := make([]byte, 1024)
		_, err = conn.Read(readByte)
		if err != nil {
			fmt.Errorf("read byte err : %+v", err)
		}
		fmt.Println(string(readByte))
	}(conn)

}
