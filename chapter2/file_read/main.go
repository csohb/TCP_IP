package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	buf, err := os.ReadFile("chapter2/file_write/Aespa")
	if err != nil {
		fmt.Errorf("os.ReadFile err : %+v", err)
	}
	fmt.Println("buf : ", buf)
	fmt.Println("string : ", string(buf))

	buf2, err := os.ReadFile("chapter3/file_write/Aespa")
	if err != nil {
		fmt.Errorf("os.ReadFile err : %+v", err)
	}
	fmt.Println("buf : ", buf2)
	fmt.Println("string : ", string(buf2))

	buf3, err := ioutil.ReadFile("chapter2/file_write/Aespa")
	if err != nil {
		fmt.Errorf("os.ReadFile err : %+v", err)
	}
	fmt.Println(string(buf3))

	buf4, err := ioutil.ReadFile("chapter3/file_write/Aespa")
	if err != nil {
		fmt.Errorf("os.ReadFile err : %+v", err)
	}
	fmt.Println(string(buf4))
}
