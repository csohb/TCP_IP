package main

import (
	"fmt"
	"log"
	"os"
	"syscall"
)

func main() {
	f, err := os.OpenFile("Aespa", os.O_RDONLY, 0755)
	if err != nil {
		log.Fatal(err)
	}
	fd := f.Fd()
	fmt.Println("file descriptor : ", fd)

	file := os.NewFile(fd, "savage file")
	data := make([]byte, 4)
	var dataAt int64 = 2
	n, err := file.ReadAt(data, dataAt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n)
	fmt.Println(data)
	fmt.Println(string(data))

	if err = syscall.Close(int(fd)); err != nil {
		log.Fatal(err)
	}
}
