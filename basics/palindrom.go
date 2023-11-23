package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func isPalindrom(phrase string) bool {
	phrase = strings.ToLower(phrase)
	phrase = removePunctuation(phrase)

	reversed := reverseString(phrase)
	return phrase == reversed
}

func reverseString(s string) string {
	runes := []rune(s)
	length := len(runes)

	for i, j := 0, length-1; i < length/2; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func removePunctuation(s string) string {
	var result []rune

	runes := []rune(s)
	for _, r := range runes {
		if unicode.IsLetter(r) {
			result = append(result, r)
		}
	}
	return string(result)
}

func main() {
	fmt.Println("Введите слово или фразу ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	phrase := scanner.Text()

	if isPalindrom(phrase) {
		fmt.Printf("\"%s\" is palindrom", phrase)
	} else {
		fmt.Printf("\"%s\" is not palindrom", phrase)
	}
}
