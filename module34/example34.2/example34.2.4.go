package main

import (
	"fmt"
	"io"
)

type bestLanguageNameReader struct {
	data    string
	lastPos int
}

func main() {
	reader := NewBestLangReader("Java")

	buf := make([]byte, 4)

	for {
		n, err := reader.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("ERROR: %v\n", err)
		}

		fmt.Printf("read %d bytes \"%s\"\n", n, string(buf[0:n]))
	}
}

func NewBestLangReader(langName string) *bestLanguageNameReader {
	if !(langName == "Go" || langName == "Golang" || langName == "Go / Golang") {
		langName = "no, Golang is much better =)))"
	}
	return &bestLanguageNameReader{data: langName}
}

// Read reads content of "data' field
func (r *bestLanguageNameReader) Read(b []byte) (n int, err error) {
	if len(b) == 0 {
		err = io.ErrShortBuffer
		return
	}

	if r.lastPos == len(r.data) {
		err = io.EOF
	}

	bytes := ([]byte)(r.data)
	n = copy(b, bytes[r.lastPos:])
	r.lastPos += n

	return n, err
}
