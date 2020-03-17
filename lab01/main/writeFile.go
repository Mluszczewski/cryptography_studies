package main

import (
	"bufio"
	"fmt"
	"os"
)

func writeEncryptedFile(encrypted string) error {

	file, err := os.Create("crypto.txt")
	if err != nil {
		fmt.Println("err")
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
		fmt.Println("err")
	}

	_, err = file.WriteString(decrypted)
	if err != nil {
		fmt.Println("Can't write into file")
	}
	file.Close()
	return file.Sync()
}

func writeListOfDecryptedFile(decrypted []string) error {
	file, err := os.Create("decrypt.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, text := range decrypted {
		fmt.Fprintln(w, text)

	}
	return w.Flush()
}

func writeKey(key int) error {
	file, err := os.Create("key-new.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("%d", key))
	if err != nil {
		return err
	}

	return file.Sync()
}
