package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"math/big"
	"os"
)

func readElgamal(param1, param2 int) string {
	var oneLine string
	file, err := os.Open("elgamal.txt")

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

func readPublicKey() *PublicKey {

	var oneLine string
	var twoLine string
	var threeLine string
	file, err := os.Open("public.txt")

	if err != nil {

	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	file.Close()

	for _, eachline := range txtlines[0:1] {
		oneLine = eachline
	}
	var gFromFile, _ = new(big.Int).SetString(oneLine, 0)
	for _, eachline := range txtlines[1:2] {
		twoLine = eachline
	}
	var pFromFile, _ = new(big.Int).SetString(twoLine, 0)
	for _, eachline := range txtlines[2:3] {
		threeLine = eachline
	}
	var yFromFile, _ = new(big.Int).SetString(threeLine, 0)

	return &PublicKey{G: gFromFile, P: pFromFile, Y: yFromFile}
}

func readPrivateKey() *PrivateKey {

	var oneLine, twoLine, threeLine, fourline string

	file, err := os.Open("private.txt")

	if err != nil {

	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	file.Close()

	for _, eachline := range txtlines[0:1] {
		oneLine = eachline
	}
	var gFromFile, _ = new(big.Int).SetString(oneLine, 0)

	for _, eachline := range txtlines[1:2] {
		twoLine = eachline
	}
	var pFromFile, _ = new(big.Int).SetString(twoLine, 0)

	for _, eachline := range txtlines[2:3] {
		threeLine = eachline
	}
	var yFromFile, _ = new(big.Int).SetString(threeLine, 0)

	for _, eachline := range txtlines[3:4] {
		fourline = eachline
	}
	var xFromFile, _ = new(big.Int).SetString(fourline, 0)

	publicKey := PublicKey{G: gFromFile, P: pFromFile, Y: yFromFile}

	return &PrivateKey{PublicKey: publicKey, X: xFromFile}
}

func readCipher() (c1, c2 *big.Int) {

	var oneLine, twoLine string

	file, err := os.Open("crypto.txt")

	if err != nil {

	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	file.Close()

	for _, eachline := range txtlines[0:1] {
		oneLine = eachline
	}
	c1, _ = new(big.Int).SetString(oneLine, 0)

	for _, eachline := range txtlines[1:2] {
		twoLine = eachline
	}
	c2, _ = new(big.Int).SetString(twoLine, 0)

	return c1, c2
}

func readFile(fileName string) []byte {

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	return b
}
