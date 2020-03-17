package main

import (
	"crypto/des"
	"errors"
	"fmt"
)

func ebcCrypto(origData, key []byte) []byte {
	tkey := make([]byte, 24, 24)
	copy(tkey, key)

	k1 := tkey[:8]
	k2 := tkey[16:]

	buf1 := encryptEbc(origData, k1)
	out := encryptEbc(buf1, k2)

	return out
}

func encryptEbc(origData, key []byte) []byte {
	if len(origData) < 1 || len(key) < 1 {
		fmt.Println(errors.New("wrong data or key"))
	}
	block, err := des.NewCipher(key)
	if err != nil {
		fmt.Println(err)
	}
	bs := 14

	for len(origData) > 0 {
		if len(origData)%bs != 0 {
			origData = append(origData, 255)
		} else {
			break
		}
	}

	out := make([]byte, len(origData))
	dst := out
	for len(origData) > 0 {
		block.Encrypt(dst, origData[:bs])
		origData = origData[bs:]
		dst = dst[bs:]
	}
	return out
}
