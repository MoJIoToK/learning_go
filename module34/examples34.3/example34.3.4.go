package main

import (
	"bufio"
	"os"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)

	writer.Write([]byte("Hello"))
	writer.Flush()
	writer.Write([]byte(", "))
	writer.Write([]byte("World"))
	writer.Write([]byte("!!!")) // break
	writer.Flush()
}
