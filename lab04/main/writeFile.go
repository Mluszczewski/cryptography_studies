package main

import (
	"bytes"
	"fmt"
	"image"
	"os"

	"golang.org/x/image/bmp"
)

func writeFile(byt []byte, fileName string) {
	image.RegisterFormat("bmp", "bmp", bmp.Decode, bmp.DecodeConfig)

	img, _, err := image.Decode(bytes.NewReader(byt))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	out, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = bmp.Encode(out, img)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	out.Close()
}
