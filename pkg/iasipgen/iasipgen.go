package iasipgen

import (
	"errors"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"image"
	"io/ioutil"
	"strings"
)

const (
	width        = 1920
	height       = 1080
	fontsize     = 120
	linepadding  = 32
	dpi          = 72
	maxLineWidth = 1700
)

var textile *truetype.Font

// LoadFont loads the Textile font for use in the generator.
// This must be done before Generate.
func LoadFont(path string) error {
	// ensure textile.ttf is available
	fontfile, err := ioutil.ReadFile(path)
	if err != nil {
		return errors.New("Couldn't load font: " + err.Error())
	}
	textile, err = freetype.ParseFont(fontfile)
	if err != nil {
		return errors.New("Couldn't parse font: " + err.Error())
	}
	if textile == nil {
		return errors.New("Couldn't parse font: font still appears to be nil")
	}
	return nil
}

// Generate generates a title card similar to the ones from It's Always Sunny in Philadelphia.
// LoadFont must be called before Generate or else the function will error out.
func Generate(title string) (img *image.RGBA, err error) {
	// create the image
	img = image.NewRGBA(image.Rect(0, 0, width, height))
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
		Face: truetype.NewFace(textile, &truetype.Options{
			Size:    fontsize,
			DPI:     dpi,
			Hinting: font.HintingFull,
		}),
	}
	title = strings.Replace(title, "\n", " ", -1)
	// determine where the line breaks in the text need to be
	words := strings.Split(title, " ")
	textsplit := []string{}
	i := 0
	for _, word := range words {
		if len(textsplit) == 0 {
			textsplit = append(textsplit, word)
			continue
		}
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
	for line := 0; line < lines; line++ {
		linelength := d.MeasureString(textsplit[line]).Round()
		y := ((height - ((fontsize + linepadding) * lines)) / 2) + ((fontsize + linepadding) * (line + 1)) - linepadding
		d.Dot = freetype.Pt((width-linelength)/2, y)
		d.DrawString(textsplit[line])
	}
	return img, nil
}
