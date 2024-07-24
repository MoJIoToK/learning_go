package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	// fpath, err := getFilePath()
	// if err != nil {
	// 	panic(err)
	// }
	fpath := "C:/Users/Nick/Desktop"

	files, err := ioutil.ReadDir(fpath)
	if err != nil {
		panic(err)
	}

	file, err := os.OpenFile("LS.txt", os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}

	writer := bufio.NewWriter(file)

	for _, info := range files {
		fType := "FILE"
		if info.IsDir() {
			fType = "DIRECTORY"
		}

		inf := fmt.Sprintf("%v (%v bytes) [%v]\n",
			info.Name(), info.Size(), fType)

		_, err := writer.WriteString(inf)
		if err != nil {
			panic(err)
		}
	}

	if err := writer.Flush(); err != nil {
		panic(err)
	}

}

/*
func getFilePath() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите строку данных: ")
	str, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	str = strings.TrimSpace(str)

	return str, nil
}
*/
