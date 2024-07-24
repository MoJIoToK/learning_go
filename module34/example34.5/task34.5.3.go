package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

const (
	minLength = 8
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите пароль " +
		"(должен содержать строчные и прописные буквы, " +
		"цифры, по крайней мере один спец.символ " +
		"и быть длинной не менее 8 символов): ")
	str, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	// удаляем символ перевода на новую строку
	str = strings.TrimSpace(str)

	// проверка длины пароля
	lenRegex := regexp.MustCompile(fmt.Sprintf(`^.{%d,}$`, minLength))
	if !lenRegex.MatchString(str) {
		fmt.Printf("Ошибка! Длина пароля менее %d\n", minLength)
		return
	}

	// проверка наличия обязательных символов
	if !regexp.MustCompile(`[a-z]+`).MatchString(str) {
		fmt.Println("Ошибка! Пароль должен содержать строчные буквы")
		return
	}
	if !regexp.MustCompile(`[A-Z]+`).MatchString(str) {
		fmt.Println("Ошибка! Пароль должен содержать прописные буквы")
		return
	}
	if !regexp.MustCompile(`[0-9]+`).MatchString(str) {
		fmt.Println("Ошибка! Пароль должен содержать цифры")
		return
	}

	mostPopularPassword := []string{
		"Qq123456",
		"Qwerty123",
	}
	join := strings.Join(mostPopularPassword, "|")
	weakPassRegex := regexp.MustCompile(fmt.Sprintf("^(%s)$", join))
	if weakPassRegex.MatchString(str) {
		fmt.Println("Предупреждение! Очень слабый пароль, придумайте другой")
		return
	}

	specCharRegex := regexp.MustCompile(`[!@#$%^&*()\-=+,./\\]+`)
	if !specCharRegex.MatchString(str) {
		fmt.Println("Ошибка! Пароль должен содержать спец.символ")
		return
	}

	fmt.Println("Введен корректный пароль")
}
