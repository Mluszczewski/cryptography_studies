package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

func cbcCrypto(origData, key []byte) []byte {

	bs := 16

	for len(origData) > 0 {
		if len(origData)%bs != 0 {
			origData = append(origData, 255)
		} else {
			break
		}
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
	}

	cipherText := make([]byte, bs+len(origData))
	iv := cipherText[:bs]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		fmt.Println(err)
	}

	cbc := cipher.NewCBCEncrypter(block, iv)
	cbc.CryptBlocks(cipherText[bs:], origData)
	return cipherText
}
