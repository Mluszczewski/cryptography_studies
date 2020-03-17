package main

import (
	"fmt"
	"log"
	"os"
)

func appendStringToCrypto(text1 []string) error {
	f, err := os.Create("crypto.txt")
	if err != nil {
		return err
	}

	for i := 0; i < len(text1); i++ {
		_, err = f.WriteString(text1[i] + "\n")
		if err != nil {
			return err
		}
	}
	f.Close()
	return f.Sync()
}

func appendStringToDecrypt(text1 []string) error {
	f, err := os.Create("decrypt.txt")
	if err != nil {
		return err
	}
	for i := 0; i < len(text1); i++ {
		_, err = f.WriteString(text1[i] + "\n")
		if err != nil {
			return err
		}
	}
	f.Close()
	return f.Sync()
}

func appendStringTKey(text1 string) error {
	f, err := os.Create("new-key.txt")
	if err != nil {
		return err
	}

	_, err = f.WriteString(text1)
	if err != nil {
		return err
	}

	f.Close()
	return f.Sync()
}

func writePreparedText(preparedText string) error {
	file, err := os.Create("plain.txt")
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

func key() string {
	key := "AbcdefghijkaBCDEfkbmhipoerWQEvcxmkFDSOEp"
	return key
}

func cipherText() string {
	cipherText := "AbcdefghijkaBCDEfkbmhipoerWQEvcxmkFDSOEp"
	return cipherText
}
