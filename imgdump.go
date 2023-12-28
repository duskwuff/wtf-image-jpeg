package main

import (
	"fmt"
	"image"
	"os"

	_ "image/jpeg"
	_ "image/png"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	b := img.Bounds()
	for y := 0; y < b.Max.Y; y++ {
		for x := 0; x < b.Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			fmt.Printf("%d,%d = %04x %04x %04x\n", x, y, r, g, b)
		}
	}
}
