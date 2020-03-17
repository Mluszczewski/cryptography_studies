package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
)

func stringToLines(fileName string, param1, param2 int) string {
	var oneLine string
	file, err := os.Open(fileName)

	if err != nil {

	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	file.Close()

	for _, eachline := range txtlines[param1:param2] {
		oneLine = eachline
	}
	return oneLine
}

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
