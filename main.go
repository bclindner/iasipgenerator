package main

import (
	"fmt"
	"github.com/bclindner/iasipgenerator/iasipgen"
	"image/jpeg"
	"os"
)

func main() {
	// make sure there's more than one argument
	if len(os.Args) < 2 {
		fmt.Println("No text specified, exiting.")
		os.Exit(1)
	}
	// open file to write image to
	file, err := os.OpenFile("out.jpg", os.O_WRONLY|os.O_CREATE, 0655)
	if err != nil {
		fmt.Println("Couldn't open output file:", err)
		os.Exit(1)
	}
	err = iasipgen.LoadFont("txtile.ttf")
	if err != nil {
		fmt.Println("Couldn't load font:", err)
		os.Exit(1)
	}
	img, err := iasipgen.Generate(os.Args[1])
	if err != nil {
		fmt.Println("Couldn't generate image:", err)
		os.Exit(1)
	}
	// write image to disk
	err = jpeg.Encode(file, img, &jpeg.Options{
		Quality: 100,
	})
	if err != nil {
		fmt.Println("Couldn't encode JPEG:", err)
		os.Exit(1)
	}
}
