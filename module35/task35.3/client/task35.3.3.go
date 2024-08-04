package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	// Подключение к сетевой службе.
	conn, err := net.Dial("tcp4", "localhost:12345")
	if err != nil {
		log.Fatal(err)
	}
	// Не забываем закрыть ресурс.
	defer conn.Close()

	// Запись запроса.
	// Не забываем добавить разделитель - перевод строки.
	_, err = conn.Write([]byte("time\n"))
	if err != nil {
		log.Fatal(err)
	}

	// Буфер для чтения данных из соединения.
	reader := bufio.NewReader(conn)
	// Считывание массива байт до перевода строки.
	b, err := reader.ReadBytes('\n')
	if err != nil {
		log.Fatal(err)
	}

	// Обработка ответа.
	fmt.Println("Ответ от сервера:", string(b))
}
