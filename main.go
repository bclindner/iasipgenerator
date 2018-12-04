package main

import (
	"fmt"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"image"
	"image/png"
	"io/ioutil"
	"os"
	"strings"
)

func centered(size int) int { return (1920 - size) / 2 }

const (
	width        = 1920
	height       = 1080
	fontsize     = 120
	linepadding  = 32
	dpi          = 72
	maxLineWidth = 1800 // anything higher results in weirdness
)

func main() {
	// make sure there's more than one argument
	if len(os.Args) < 2 {
		fmt.Println("No text specified, exiting.")
		os.Exit(1)
	}
	text := os.Args[1]
	// open file to write image to
	file, err := os.OpenFile("out.png", os.O_WRONLY|os.O_CREATE, 0655)
	if err != nil {
		fmt.Println("Couldn't open output file")
		os.Exit(1)
	}
	// open Textile font
	fontfile, err := ioutil.ReadFile("textile.ttf")
	if err != nil {
		fmt.Println("Couldn't find textile.ttf")
		os.Exit(1)
	}
	f, err := freetype.ParseFont(fontfile)
	defer file.Close()
	// create the image
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	// fill image with black
	for x := 0; x < img.Bounds().Dx(); x++ {
		for y := 0; y < img.Bounds().Dy(); y++ {
			img.Set(x, y, image.Black)
		}
	}
	// instantiate text drawer
	d := &font.Drawer{
		Dst: img,
		Src: image.White,
		Face: truetype.NewFace(f, &truetype.Options{
			Size:    fontsize,
			DPI:     dpi,
			Hinting: font.HintingFull,
		}),
	}
	// determine where the line breaks in the text need to be
	textsplit := []string{""}
	i := 0
	for _, word := range strings.Split(text, " ") {
		linelength := d.MeasureString(textsplit[i] + word).Round()
		if linelength >= maxLineWidth {
			i++
			textsplit = append(textsplit, word)
		} else {
			textsplit[i] += " " + word
		}
	}
	// write the text to the image
	lines := len(textsplit)
	fmt.Printf("Split into %d lines\n", lines)
	for line := 0; line < lines; line++ {
		linelength := d.MeasureString(textsplit[line]).Round()
		y := ((height - ((fontsize + linepadding) * lines)) / 2) + ((fontsize + linepadding) * (line + 1)) - linepadding
		d.Dot = freetype.Pt((width-linelength)/2, y)
		d.DrawString(textsplit[line])
	}
	// write image to disk
	err = png.Encode(file, img)
	if err != nil {
		panic(err)
	}
}
