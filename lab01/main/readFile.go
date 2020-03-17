package main

import (
	"bufio"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func readFile(fileName string) string {

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}

func readKey(r io.Reader) string {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)

	var result []byte

	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, byte(x))
	}

	return string(result[1])
}
