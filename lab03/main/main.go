package main

import (
	"flag"
	"strings"
)

func main() {
	flag.Bool("p", false, "Prepare text to encoding")
	flag.Bool("e", false, "Encryption")
	flag.Bool("k", false, "Cryptoanalysis based solely on the cryptogram")
	flag.Parse()

	flagset := make(map[string]bool)
	flag.Visit(func(f *flag.Flag) { flagset[f.Name] = true })

	if flagset["p"] {
		prepare := readFile("orig.txt")
		prepareText := prepareText(prepare)
		writePreparedText(prepareText)
	}

	var zbiorCrypted []string
	if flagset["e"] {
		var cipher string
		readFile("key.txt")
		j := 1
		for i := 0; i < 20; i++ {
			cipher = encryptIt(stringToLines("plain.txt", i, j), key())
			zbiorCrypted = append(zbiorCrypted, cipher)
			j++
			if j == 21 {
				break
			}
			appendStringToCrypto(zbiorCrypted)
		}
	}

	if flagset["k"] {
		var zbiorDecrypted []string
		var decipher string

		w := 1
		for i := 0; i < 21; i++ {
			decipher = decryptIt(stringToLines("crypto.txt", i, w), cipherText())
			zbiorDecrypted = append(zbiorDecrypted, decipher)
			w++
			if w == 21 {
				break
			}
			enc := strings.Replace(decipher, " ", "", -1)
			txt := make([]int, len(enc))
			for i := 0; i < len(txt); i++ {
				txt[i] = int(enc[i] % 26)
			}
			bestFit, _ := 1e100, ""
			for j := 1; j <= 26; j++ {
				key := make([]byte, j)
				fit := freqEveryNth(txt, key)
				sKey := string(key)
				if fit < bestFit {
					bestFit, _ = fit, sKey
				}
			}
		}
		appendStringToDecrypt(zbiorDecrypted)
	}
}
