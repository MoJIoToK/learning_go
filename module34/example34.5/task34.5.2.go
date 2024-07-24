package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите email: ")
	str, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	// удаляем символ перевода на новую строку
	str = strings.TrimSpace(str)

	emailRegex := regexp.MustCompile(`^[a-zA-Z][\w.-]+@[\w.-]+$`)
	isMatch := emailRegex.MatchString(str)

	fmt.Printf("Ввод \"%s\", совпадение с шаблоном: %v", str, isMatch)
}
