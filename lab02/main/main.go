package main

import (
	"flag"
	"strings"
)

func main() {

	flag.Bool("p", false, "Prepare text to encoding")
	flag.Bool("e", false, "Encryption")
	flag.Bool("d", false, "Decryption")
	flag.Bool("k", false, "Cryptoanalysis based solely on the cryptogram")
	flag.Parse()

	flagset := make(map[string]bool)
	flag.Visit(func(f *flag.Flag) { flagset[f.Name] = true })

	if flagset["p"] {
		prepare := readFile("orig.txt")
		prepareText := prepareText(prepare)
		writePreparedText(prepareText)
	}

	if flagset["e"] {
		v := generateKey(readFile("key.txt"))
		ct := v.encryptCipher(readFile("orig.txt"))
		writeEncryptedFile(ct)
	}

	if flagset["d"] {
		v := generateKey(readFile("key.txt"))
		ct := v.encryptCipher(readFile("orig.txt"))
		dt := v.decryptCipher(ct)
		writeDecryptedFile(dt)
	}

	if flagset["k"] {
		var encoded = readFile("crypto.txt")
		enc := strings.Replace(encoded, " ", "", -1)
		txt := make([]int, len(enc))
		for i := 0; i < len(txt); i++ {
			txt[i] = int(enc[i] - 'A')
		}
		bestFit, bestKey := 1e100, ""
		for j := 1; j <= 26; j++ {
			key := make([]byte, j)
			fit := freqEveryNth(txt, key)
			sKey := string(key)
			if fit < bestFit {
				bestFit, bestKey = fit, sKey
			}
		}
		writeKey(bestKey)
		writeDecryptedFile(decrypt(enc, bestKey))
	}
}
