package main

import (
	"fmt"
	"log"
	"net/rpc"

	"module35/task35.8.3/model"
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

	var resp float64

	point1 := model.Point{X: 1, Y: 1}
	point2 := model.Point{X: 4, Y: 5}
	points := model.Points{A: point1, B: point2}

	// Удаленный вызов процедуры Server.Time. Должна быть ошибка.
	err = client.Call("Server.Dist", points, &resp)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("time: ", resp)

	//// Удаленный вызов процедуры Server.Time. Должен быть ответ.
	//err = client.Call("Server.Time", "time", &resp)
	//if err != nil {
	//	fmt.Println("ошибка:", err)
	//}
	//fmt.Println("time:", resp)
}
