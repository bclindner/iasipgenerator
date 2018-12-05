package main

import (
	"fmt"
	"github.com/bclindner/iasipgenerator/iasipgen"
	"image/png"
	"os"
)

func main() {
	// make sure there's more than one argument
	if len(os.Args) < 2 {
		fmt.Println("No text specified, exiting.")
		os.Exit(1)
	}
	// open file to write image to
	file, err := os.OpenFile("out.png", os.O_WRONLY|os.O_CREATE, 0655)
	if err != nil {
		fmt.Println("Couldn't open output file")
		os.Exit(1)
	}
	img, err := iasipgen.Generate(os.Args[1])
	// write image to disk
	err = png.Encode(file, img)
	if err != nil {
		panic(err)
	}
}
