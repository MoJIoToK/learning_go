package main

import (
	"log"
	"math/rand"
	"net"
	"task35.8.1/model"
	"time"
)

//Server - сервер для сетевой службы который составленный согласно заданию 35.8.1.

// Константы сетевого адреса который будет прослушивать сетевая служба.
// Данные константы вставляются в функции стандартного пакета `net`.
const (
	//Протокол сетевой службы
	network = "tcp4"

	//Порт для прослушивания(адрес конкретного приложения, к которому стоит обращаться)
	addr = ":12345"
)

func main() {

	//Создание и запуск сетевой службы с помощью функции `net.Listener()`. Функция возвращает экземпляр
	//объекта `Listener`.
	listener, err := net.Listen(network, addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	//Приём входящих подключений в бесконечном цикле, для обеспечений беспрерывного
	//обслуживания подключения.
	for {
		//Приём подключения с помощью функции `net.Accept()`. Данная функция возвращает экземляр
		//объекта подключения - `Conn`.
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		//Вызов обработчика в отдельной горутине для обеспечения множественной одновременной обработки
		//разных подключений.
		go handleAphorism(conn)

	}

}

// handleAphorism - функция обработки подключения. Функция отправляет каждому подключению
// случайную поговорку с го-сайта раз в три секунды.
func handleAphorism(conn net.Conn) {

	for {

		//реализация выбора рандомного сообщения из списка цитат.
		msg := model.Aphorisms[rand.Intn(len(model.Aphorisms))]

		//Реализация таймера с помощью стандартного пакета `time`
		t := time.NewTimer(3 * time.Second)
		<-t.C
		conn.Write([]byte("3 sec message: "))
		conn.Write([]byte(msg + "\n\r"))
	}

}
