package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("./hi.txt")
	if err != nil {
		panic(err)
	}
	buf := make([]byte, 10)

	for {
		n, err := f.Read(buf)
		if err == io.EOF {
			break
		}

		fmt.Println("> ", string(buf[:n]))

	}
}
