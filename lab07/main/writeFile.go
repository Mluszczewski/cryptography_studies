package main

import "os"

import "fmt"

import "math/big"

func writePrivateKey(privateKey *PrivateKey) error {
	f, err := os.Create("private.txt")
	if err != nil {
		return err
	}

	_, err = f.WriteString(fmt.Sprintln(privateKey.G))
	_, err = f.WriteString(fmt.Sprintln(privateKey.P))
	_, err = f.WriteString(fmt.Sprintln(privateKey.Y))
	_, err = f.WriteString(fmt.Sprintln(privateKey.X))

	if err != nil {
		return err
	}

	f.Close()
	return f.Sync()
}

func writePublicKey(publicKey *PublicKey) error {
	f, err := os.Create("public.txt")
	if err != nil {
		return err
	}

	_, err = f.WriteString(fmt.Sprintln(publicKey.G))
	_, err = f.WriteString(fmt.Sprintln(publicKey.P))
	_, err = f.WriteString(fmt.Sprintln(publicKey.Y))

	if err != nil {
		return err
	}

	f.Close()
	return f.Sync()
}

func writeCipher(c1, c2 *big.Int) error {
	f, err := os.Create("crypto.txt")
	if err != nil {
		return err
	}

	_, err = f.WriteString(fmt.Sprintln(c1))
	_, err = f.WriteString(fmt.Sprintln(c2))

	if err != nil {
		return err
	}

	f.Close()
	return f.Sync()
}

func writeDecrypt(encrypted []byte) error {

	file, err := os.Create("decrypt.txt")
	if err != nil {
		fmt.Println("err")
	}

	_, err = file.Write(encrypted)
	if err != nil {
		fmt.Println("Can't write into file")
	}
	file.Close()
	return file.Sync()
}
