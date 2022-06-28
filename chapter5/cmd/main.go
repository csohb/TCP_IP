package main

import "fmt"

func main() {
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
}
