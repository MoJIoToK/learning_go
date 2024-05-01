package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите данные: ")
	data, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Вы ввели данные: %v\n", data)
}
