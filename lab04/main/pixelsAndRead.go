package main

import (
	"errors"
	"fmt"
	"image"

	"io"

	"os"

	"golang.org/x/image/bmp"
)

type pixel struct {
	R int
	G int
	B int
	A int
}

var errUnsuported = errors.New("bmp: unsupported BMP image")

func readFileAndParse(myImage string) [][]pixel {
	image.RegisterFormat("bmp", "bmp", bmp.Decode, bmp.DecodeConfig)

	file, err := os.Open(myImage)

	if err != nil {
		os.Exit(1)
	}

	defer file.Close()

	pixels, err := getPixels(file)

	if err != nil {
		fmt.Println("Error: Image could not be decoded")
		os.Exit(1)
	}

	return pixels

}

func getPixels(file io.Reader) ([][]pixel, error) {
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	var pixels [][]pixel
	for y := 0; y < height; y++ {
		var row []pixel
		for x := 0; x < width; x++ {
			row = append(row, rgbaToPixel(img.At(x, y).RGBA()))
		}
		pixels = append(pixels, row)
	}
	return pixels, nil
}

func rgbaToPixel(r, g, b, a uint32) pixel {
	return pixel{int(r / 257), int(g / 257), int(b / 257), int(a / 257)}
}
