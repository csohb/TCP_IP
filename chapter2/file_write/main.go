package main

import (
	"fmt"
	"os"
)

func main() {
	var buf []byte
	buf = []byte("I'm on the Next Level")

	err := os.WriteFile("chapter2/file_write/Aespa", buf, 0644)
	if err != nil {
		fmt.Errorf("os.WriteFile err : %+v", err)
	}
}
