package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	//Ввод наименования документа с входными данными
	fmt.Print("Enter input file name: ")
	inputFileName, err := scanFileName()
	if err != nil {
		panic(err)
	}
	fmt.Println("OK. Input file name is writing - ", inputFileName)

	//Ввод наименования документа для записи результатов
	fmt.Print("Enter output file name: ")
	outputFileName, err := scanFileName()
	if err != nil {
		panic(err)
	}
	fmt.Println("OK. Output file name is writing - ", outputFileName)

	//Чтение файла с входными данными и вызов дальнейших функций приложения
	if err := readWriteOperation(inputFileName, outputFileName); err != nil {
		panic(err)
	}

}

// scanFileName - функция для чтения наименований файлов. На входе функция ничего не принимает.
// Функция возвращает строку и ошибку.
func scanFileName() (str string, err error) {
	reader := bufio.NewReader(os.Stdin)
	str, err = reader.ReadString('\n')
	if err != nil {
		return "", errors.New("Ошибка чтения ввода")
	}

	str = strings.TrimSpace(str)
	return str, nil
}

// readWriteOperation - основная функция приложения, которая отвечает за чтение файла со входными данными,
// вызов функций для обработки данных, и запись обработанных данных в файл.
// Входными параметрами функции являются наименования файлов. Функция возвращает ошибку.
func readWriteOperation(inputFileName, outputFileName string) error {
	//Открытие/создание файла для записи результирующих данных
	fileWrit, err := os.OpenFile(outputFileName, os.O_CREATE|os.O_APPEND|os.O_RDWR|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	//Открытие файла для чтения входных данных
	file, err := os.OpenFile(inputFileName, os.O_RDONLY, 0777)
	if err != nil {
		panic(err)
	}

	//Создание экземпляров Writer и Reader
	fileWriter := bufio.NewWriter(fileWrit)
	fileReader := bufio.NewReader(file)

	//Построчное чтение файла до его конца. Каждая прочитанная строка обрабатывается в обработчике и возвращается
	//как результат, который затем пишется в файл с помощью буферизации.
	for {
		line, _, err := fileReader.ReadLine()
		if err == io.EOF {
			break
		}

		//вызов обработчика строки
		res, err := handler(string(line))
		if err != nil {
			log.Println(err)
		}
		fileWriter.WriteString(res)
	}

	fileWriter.Flush()
	fileWrit.Close()

	file.Close()

	return nil
}

// handler - функция обрабатывает строку путем её сопоставления с регулярным выражением. Регулярное выражение
// построено таким образом, чтобы строка делилась на группы захвата - два операнда и оператор между ними.
// Операнды могут быть как положительными числами, так и отрицательными. Оператор может располагаться с пробелами
// между операндами. Функция возвращает строку с результатом и ошибку.
func handler(s string) (string, error) {
	re := regexp.MustCompile(`\s*([-+]?\d*\.?\d+)\s*([\+\-\*/])\s*([-+]?\d*\.?\d+)\s*`)
	match := re.FindStringSubmatch(s)

	if len(match) != 4 {
		er := fmt.Sprintf("%v - Некорректное выражение", match)
		return "", errors.New(er)
	}

	operand1, err := parser(match[1])
	if err != nil {
		return "", err
	}

	operator := match[2]

	operand2, err := parser(match[3])
	if err != nil {
		return "", err
	}

	res, err := eval(operand1, operator, operand2)
	if err != nil {
		return "", err
	}

	resStr := fmt.Sprintf("%v%v%v=%v\n", operand1, operator, operand2, res)

	return resStr, nil

}

// parser - функция преобразует строковое представление числа в формат float64. На вход функция принимает - строку.
// На выходе функция возвращает число в формате float64 и ошибку error.
func parser(str string) (float64, error) {
	par, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0, err
	}
	return par, nil
}

// eval - функция для совершения математических операций над числами. На вход подается два операнда с типом float64
// и строка с оператором. На выходе получается результат типа float64 и ошибка.
func eval(operand1 float64, operator string, operand2 float64) (float64, error) {
	switch operator {
	case "+":
		return operand1 + operand2, nil
	case "-":
		return operand1 - operand2, nil
	case "*":
		return operand1 * operand2, nil
	case "/":
		if operand2 == 0 {
			return 0, errors.New("Деление на ноль")
		}
		return operand1 / operand2, nil
	default:
		return 0, errors.New("Неизвестный оператор")
	}
}
