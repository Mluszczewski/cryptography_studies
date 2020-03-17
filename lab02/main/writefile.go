package main

import (
	"fmt"
	"log"
	"os"
)

func writeEncryptedFile(encrypted string) error {

	file, err := os.Create("crypto.txt")
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.WriteString(encrypted)
	if err != nil {
		fmt.Println("Can't write into file")
	}
	file.Close()
	return file.Sync()
}

func writeDecryptedFile(decrypted string) error {
	file, err := os.Create("decrypt.txt")
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.WriteString(decrypted)
	if err != nil {
		fmt.Println("Can't write into file")
	}
	file.Close()
	return file.Sync()
}

func writePreparedText(preparedText string) error {
	file, err := os.Create("orig.txt")
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.WriteString(preparedText)
	if err != nil {
		fmt.Println("Can't write into file")
	}
	file.Close()
	return file.Sync()
}

func writeKey(key string) error {
	file, err := os.Create("key-crypto.txt")
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.WriteString(key)
	if err != nil {
		fmt.Println("Can't write into file")
	}
	file.Close()
	return file.Sync()
}
