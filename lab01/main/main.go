package main

import (
	"flag"
	"fmt"
	"strings"
)

var msg = readFile("plain.txt")

func main() {
	flag.Bool("c", false, "Caesar Cipher")
	flag.Bool("a", false, "Affine Ciher")
	flag.Bool("e", false, "Encryption")
	flag.Bool("d", false, "Decryption")
	flag.Bool("k", false, "Cryptoanalysis based solely on the cryptogram")
	flag.Bool("j", false, "Cryptoanalysis based on public text")
	flag.Parse()

	flagset := make(map[string]bool)
	flag.Visit(func(f *flag.Flag) { flagset[f.Name] = true })

	if flagset["c"] {
		if flagset["e"] {
			key := readFile("key.txt")
			shift, _ := readKey(strings.NewReader(key))
			caesarText := encryptCaesar(msg, shift)
			writeEncryptedFile(caesarText)
		} else if flagset["d"] {
			key := readFile("key.txt")
			shift, _ := readKey(strings.NewReader(key))
			caesarText := encryptCaesar(msg, shift)
			writeDecryptedFile(decryptCaesar(caesarText, shift))
		} else if flagset["k"] {
			caesarText := readFile("crypto.txt")
			writeListOfDecryptedFile(bruteForceCaesar(caesarText))
		} else if flagset["j"] {
			key := readFile("key.txt")
			extra := readFile("extra.txt")
			shift, _ := readKey(strings.NewReader(key))
			eKey := encryptCaesar(extra, shift)
			writeKey(cryptoAnalysCaesarKey(extra, eKey))
			writeDecryptedFile(decryptByCryptoAnalysCaesar(eKey, cryptoAnalysCaesarKey(extra, eKey)))
		}
	}

	if flagset["a"] {
		if flagset["e"] {
			key := readFile("key.txt")
			shift, factor := readKey(strings.NewReader(key))
			cipherText := encryptAffine(msg, shift, factor)
			writeEncryptedFile(cipherText)
		} else if flagset["d"] {
			key := readFile("key.txt")
			shift, factor := readKey(strings.NewReader(key))
			cipherText := encryptAffine(msg, shift, factor)
			writeDecryptedFile(decryptAffine(cipherText, shift, factor))
		} else if flagset["k"] {
			cipherText := readFile("crypto.txt")
			writeListOfDecryptedFile(bruteForceAffine(cipherText))
		} else if flagset["j"] {
			key := readFile("key.txt")
			shift, factor := readKey(strings.NewReader(key))
			eKey := encryptAffine(msg, shift, factor)
			fmt.Println(decryptByCryptoAnalysAffine(msg, eKey))
		}
	}
}
