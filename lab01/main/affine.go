package main

import (
	"log"
)

func encryptAffine(msg string, a, b int) string {

	cipher := []byte(msg)

	for index, byteValue := range cipher {
		if cipher[index] != ' ' {
			if byteValue >= 'a' && byteValue <= 'z' {
				cipher[index] = upperCaseAlphabet[(rune((rune(a)*(rune(byteValue)-'a')))+rune(b))%26]
			} else if byteValue >= 'A' && byteValue <= 'Z' {
				cipher[index] = upperCaseAlphabet[(rune((rune(a)*(rune(byteValue)-'A')))+rune(b))%26]
			}
		} else {
			cipher = append(cipher, cipher[index])
		}
	}
	if a%2 == 0 || a == 13 || b > 25 {
		log.Fatal("Wrong key used for Affine Cipher")
	}
	return string(cipher)
}

func decryptAffine(cipher string, a, b int) string {
	msg := []byte(cipher)
	aInv := 0
	flag := 0

	for i := 0; i < 26; i++ {
		flag = (a * i) % 26
		if flag == 1 {
			aInv = i
		}
	}

	if aInv%2 == 0 || aInv == 13 || b > 25 {
		log.Fatal("Wrong key used Affine Cipher")
	}

	for index, byteValue := range cipher {
		if cipher[index] != ' ' {
			if byteValue >= 'a' && byteValue <= 'z' {
				msg[index] = upperCaseAlphabet[(rune(aInv)*(byteValue+'a'-rune(b)))%26]
			} else if byteValue >= 'A' && byteValue <= 'Z' {
				msg[index] = upperCaseAlphabet[(rune(aInv)*(byteValue+'A'-rune(b)))%26]
			}
		} else {
			msg = append(msg, msg[index])
		}
	}

	return string(msg)
}

func bruteForceAffine(cipher string) []string {
	msg := []byte(cipher)

	iteration := []string{}

	for b := 1; b <= 25; b++ {
		for a := 1; a <= 25; a = a + 2 {
			if a == 13 {
				a = 15
			}
			aInv := 0
			flag := 0

			for i := 0; i < 26; i++ {
				flag = (a * i) % 26
				if flag == 1 {
					aInv = i
				}

			}
			for index, byteValue := range cipher {
				if cipher[index] != ' ' {
					if byteValue >= 'a' && byteValue <= 'z' {
						msg[index] = upperCaseAlphabet[(rune(aInv)*(byteValue+'a'-rune(b)))%26]
					} else if byteValue >= 'A' && byteValue <= 'Z' {
						msg[index] = upperCaseAlphabet[(rune(aInv)*(byteValue+'A'-rune(b)))%26] // a^-1 * ((a * x )- b) mod 26
					}
				} else {
					msg = append(msg, msg[index])
				}
			}
			iteration = append(iteration, string(msg))
		}
	}
	return iteration
}

func decryptByCryptoAnalysAffine(cipher string, publicKey string) []int {
	msg := []byte(cipher)
	iteration := []string{}
	var key []int
	var a int
	var b int

	for b = 1; b <= 25; b++ {
		for a = 1; a <= 25; a = a + 2 {
			if a == 13 {
				a = 15
			}
			aInv := 0
			flag := 0

			for i := 0; i < 26; i++ {
				flag = (a * i) % 26
				if flag == 1 {
					aInv = i
				}

			}
			for index, byteValue := range cipher {
				if cipher[index] != ' ' {
					if byteValue >= 'a' && byteValue <= 'z' {
						msg[index] = upperCaseAlphabet[(rune(aInv)*(byteValue+'a'-rune(b)))%26]
					} else if byteValue >= 'A' && byteValue <= 'Z' {
						msg[index] = upperCaseAlphabet[(rune(aInv)*(byteValue+'A'-rune(b)))%26] // a^-1 * ((a * x )- b) mod 26
					}
				} else {
					msg = append(msg, msg[index])
				}
			}
			iteration = append(iteration, string(msg))
			for i := range iteration {
				if iteration[i] == publicKey {
					key = append(key, a)
				}
			}
		}
	}
	return key
}
