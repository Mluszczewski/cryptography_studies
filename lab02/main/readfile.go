package main

import (
	"io/ioutil"
	"log"
	"os"
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
