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
)

func centered(size int) int { return (1920 - size) / 2 }

const (
	width       = 1920
	height      = 1080
	fontsize    = 120
	linepadding = 32
	dpi         = 72
)

func main() {
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
	// draw text, centered, on the image
	d := &font.Drawer{
		Dst: img,
		Src: image.White,
		Face: truetype.NewFace(f, &truetype.Options{
			Size:    fontsize,
			DPI:     dpi,
			Hinting: font.HintingFull,
		}),
	}
	lines := len(os.Args) - 1
	for line := 1; line < lines+1; line++ {
		linelength := d.MeasureString(os.Args[line]).Round() // rounding is fine for this
		y := ((height - ((fontsize + linepadding) * lines)) / 2) + ((fontsize + linepadding) * (line)) - linepadding
		d.Dot = freetype.Pt((width-linelength)/2, y)
		d.DrawString(os.Args[line])
		fmt.Printf("Line %d size: %d pixels\n", line, linelength)
	}
	// write image to disk
	err = png.Encode(file, img)
	if err != nil {
		panic(err)
	}
}
