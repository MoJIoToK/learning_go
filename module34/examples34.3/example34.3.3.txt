package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		buf := make([]byte, 1024)

		n1, err := conn.Read(buf)
		if err != nil {
			panic(err)
		}
		fmt.Println("Income: ", string(buf[:n1]))

		n2, err := conn.Write(buf[:n1])
		if err != nil {
			panic(err)
		}
		fmt.Println("Outcome: ", string(buf[:n2]))

		conn.Close()
	}
}
