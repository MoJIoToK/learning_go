package main

import (
	"io/ioutil"
	"os"
)

func main() {
	f, err := os.OpenFile("out.txt", os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}

	_, _ = f.Write([]byte("Hi"))
	_, _ = f.Write([]byte(" "))
	_, _ = f.Write([]byte("there"))
	_, _ = f.Write([]byte("!!!"))

	rea := ioutil.ReadAll()
}
