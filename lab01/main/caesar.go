package main

const (
	lowerCaseAlphabet = "abcdefghijklmnopqrstuvwxyz"
	upperCaseAlphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func caesarCipher(inputText string, rot int) string {
	rot %= 26
	rotatedText := []byte(inputText)

	for index, byteValue := range rotatedText {
		if byteValue >= 'a' && byteValue <= 'z' {
			rotatedText[index] = lowerCaseAlphabet[(int((26+(byteValue-'a')))+rot)%26]
		} else if byteValue >= 'A' && byteValue <= 'Z' {
			rotatedText[index] = upperCaseAlphabet[(int((26+(byteValue-'A')))+rot)%26]
		}
	}
	return string(rotatedText)
}

func bruteForceCaesar(plaintext string) []string {
	var key int
	var iteration []string

	for i := 0; i < 25; i++ {
		key++
		brute := caesarCipher(plaintext, -key)

		iteration = append(iteration, brute)
	}

	return iteration
}

func cryptoAnalysCaesarKey(publicKey string, encryptedKey string) int {
	pKey := []rune(publicKey)
	eKey := []rune(encryptedKey)
	var key int

	key = int(eKey[1]) - int(pKey[1])
	if key < 5 {
		key += 26
	}

	return key
}

func encryptCaesar(plaintext string, key int) string {
	return caesarCipher(plaintext, key)
}

func decryptCaesar(plaintext string, key int) string {
	return caesarCipher(plaintext, -key)
}

func decryptByCryptoAnalysCaesar(plaintext string, key int) string {
	return caesarCipher(plaintext, -key)
}
