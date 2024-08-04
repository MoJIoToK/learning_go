package main

import (
	"log"
	"math"
	"module35/task35.8.3/model"
	"net"
	"net/http"
	"net/rpc"
)

const (
	addr    = ":12345"
	network = "tcp4"
)

type Server int

func (s *Server) Dist(p model.Points, resp *float64) error {
	deltaX := math.Abs(p.A.X - p.B.X)
	deltaY := math.Abs(p.A.Y - p.B.Y)
	*resp = math.Sqrt(deltaX*deltaX + deltaY*deltaY)
	return nil
}

func main() {
	// Создаем указатель на переменную типа Server.
	server := new(Server)
	// Регистрируем методы типа Server в службе RPC.
	rpc.Register(server)
	// Регистрируем HTTP-обработчик для службы RPC.
	rpc.HandleHTTP()
	// Создаём сетевую службу.
	listener, err := net.Listen(network, addr)
	if err != nil {
		log.Fatal(err)
	}
	// Запускаем HTTP-сервер поверх созданной сетевой службы.
	http.Serve(listener, nil)
}
