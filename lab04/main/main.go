package main

func main() {
	pic1 := "plain.bmp"
	pic2 := "plain2.bmp"

	keyEcb := []byte("PanTadeu")
	keyCbc := []byte("0123456789abcdef")

	headerEcb1 := getHeader(pic1)
	bodyEcb1 := getBody(pic1)

	headerEcb2 := getHeader(pic2)
	bodyEcb2 := getBody(pic2)

	headerCbc1 := getHeader(pic1)
	bodyCbc1 := getBody(pic1)

	headerCbc2 := getHeader(pic2)
	bodyCbc2 := getBody(pic2)

	cipherEcb1 := encryptEbc(bodyEcb1, keyEcb)
	for i := 0; i < len(cipherEcb1); i++ {
		headerEcb1 = append(headerEcb1, cipherEcb1[i])
	}

	cipherEcb2 := encryptEbc(bodyEcb2, keyEcb)
	for i := 0; i < len(cipherEcb2); i++ {
		headerEcb2 = append(headerEcb2, cipherEcb2[i])
	}

	writeFile(headerEcb1, "ecb_crypto.bmp")
	writeFile(headerEcb2, "ecb_crypto2.bmp")

	cipherCbc1 := cbcCrypto(bodyCbc1, keyCbc)
	for i := 0; i < len(cipherCbc1); i++ {
		headerCbc1 = append(headerCbc1, cipherCbc1[i])
	}

	cipherCbc2 := cbcCrypto(bodyCbc2, keyCbc)
	for i := 0; i < len(cipherCbc2); i++ {
		headerCbc2 = append(headerCbc2, cipherCbc2[i])
	}

	writeFile(headerCbc1, "cbc_crypto.bmp")
	writeFile(headerCbc2, "cbc_crypto2.bmp")

}
