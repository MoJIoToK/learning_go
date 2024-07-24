package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func main() {
	contentBytes, err := ioutil.ReadFile("./index.html")
	if err != nil {
		panic(err)
	}

	re := regexp.MustCompile(`class="([a-zA-Z0-9_\-\s]+)"`)

	submatches := re.FindAllStringSubmatch(string(contentBytes), -1)

	for _, s := range submatches {
		classes := strings.Split(s[1], " ")
		for _, c := range classes {
			fmt.Println("Найден класс", c)
		}
	}
}
