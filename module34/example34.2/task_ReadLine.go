package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	f, _ := os.OpenFile(filename, os.O_RDONLY, 0777)

	fileReader := bufio.NewReader(f)

	for {
		line, _, err := fileReader.ReadLine()
		if err == io.EOF {
			break
		}
		fmt.Println(string(line))
	}
}
