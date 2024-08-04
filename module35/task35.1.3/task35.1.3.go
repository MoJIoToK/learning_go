package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

var UndefinedDateFormat = errors.New("undefined date format")

func main() {
	dates := []string{
		"12.09.1978",
		"1990/06/10",
		"08.03.2021",
		"12.04.1986",
		"25 dec 1988",
		"2001/05/25",
	}

	for _, d := range dates {
		year, month, day, err := parseDate(d)
		if err != nil {
			fmt.Println("ERROR!", err, "-", d)
			continue
		}

		fmt.Printf("%04d.%02d.%02d\n", year, month, day)
	}
}

func parseDate(date string) (year, month, day int64, err error) {
	re := regexp.MustCompile(`([\d]+)`)
	var dayInt, monthInt, yearInt int64

	dateSubmatch := re.FindAllStringSubmatch(date, -1)

	if len(dateSubmatch) != 3 {
		return 0, 0, 0, UndefinedDateFormat
	}

	if len(dateSubmatch[0][0]) == 2 {
		dayInt = strToInt(dateSubmatch[0][0])
		monthInt = strToInt(dateSubmatch[1][0])
		yearInt = strToInt(dateSubmatch[2][0])
	} else {
		dayInt = strToInt(dateSubmatch[2][0])
		monthInt = strToInt(dateSubmatch[1][0])
		yearInt = strToInt(dateSubmatch[0][0])
	}

	// TODO: try dd.mm.YYYY format

	// TODO: or try YYYY/mm/dd format

	// or error

	return yearInt, monthInt, dayInt, err
}

func strToInt(s string) int64 {
	n, _ := strconv.ParseInt(s, 10, 64)
	return n
}
