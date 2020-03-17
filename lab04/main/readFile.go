package main

import (
	"bufio"
	"fmt"
	"os"
)

func readAndParse(fileToUpload string) []byte {

	file, err := os.Open(fileToUpload)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	fileInfo, _ := file.Stat()
	var size int64 = fileInfo.Size()

	byt := make([]byte, size)

	buffer := bufio.NewReader(file)
	_, err = buffer.Read(byt)

	return byt
}

func getHeader(fliename string) []byte {
	bytes := readAndParse(fliename)
	var header []byte
	for i := 0; i < len(bytes); i++ {
		if i <= 40 {
			header = append(header, bytes[i])
		}
	}
	return header
}

func getBody(filename string) []byte {
	bytes := readAndParse(filename)
	var body []byte
	for i := 0; i < len(bytes); i++ {
		if i > 40 {
			body = append(body, bytes[i])
		}
	}
	return bytes
}
