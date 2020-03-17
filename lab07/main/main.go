package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
)

func fromHex(hex string) *big.Int {
	n, ok := new(big.Int).SetString(hex, 16)
	if !ok {
		panic("failed to parse hex number")
	}
	return n
}

func main() {

	flag.Bool("k", false, "generate pair of keys")
	flag.Bool("e", false, "Encryption")
	flag.Bool("d", false, "Decryption")
	flag.Parse()

	flagset := make(map[string]bool)
	flag.Visit(func(f *flag.Flag) { flagset[f.Name] = true })

	var primeHex, generatorHex string
	if flagset["k"] {
		primeHex = readElgamal(1, 2)
		generatorHex = readElgamal(0, 1)
	}
	priv := &PrivateKey{}

	if primeHex != "" && generatorHex != "" {
		priv = &PrivateKey{
			PublicKey: PublicKey{
				G: fromHex(generatorHex),
				P: fromHex(primeHex),
			},
			X: fromHex("42"),
		}
		priv.Y = new(big.Int).Exp(priv.G, priv.X, priv.P)
		writePrivateKey(priv)
		writePublicKey(&PublicKey{priv.G, priv.P, priv.Y})
	}

	if flagset["e"] {
		msg := readFile("plain.txt")
		publicKey := readPublicKey()
		c1, c2, err := encrypt(rand.Reader, publicKey, msg)
		if err != nil {
			fmt.Printf("error: %s", err)
		}
		writeCipher(c1, c2)
	}

	if flagset["d"] {
		privateKey := readPrivateKey()
		c1, c2 := readCipher()
		// fmt.Println(c1, c2)
		message, _ := decrypt(privateKey, c1, c2)
		writeDecrypt(message)
	}
}
