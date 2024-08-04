package main

import (
	"fmt"
	"log"
	"net/rpc"
)

const (
	addrClient    = ":12345"
	networkClient = "tcp4"
)

func main() {

	// Создаем клиента службы RPC.
	client, err := rpc.DialHTTP(networkClient, addrClient)
	if err != nil {
		log.Fatal(err)
	}

	var resp string

	// Удаленный вызов процедуры Server.Time. Должна быть ошибка.
	err = client.Call("Server.Time", "unknown message", &resp)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("time: ", resp)

	// Удаленный вызов процедуры Server.Time. Должен быть ответ.
	err = client.Call("Server.Time", "time", &resp)
	if err != nil {
		fmt.Println("ошибка:", err)
	}
	fmt.Println("time:", resp)
}
