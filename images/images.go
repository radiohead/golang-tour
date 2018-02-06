package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	w, h     int
	clrModel color.Model
}

func (img Image) At(x, y int) color.Color {
	w1, h1 := img.w+1, img.h+1
	x1, y1 := x+1, y+1

	return color.RGBA{
		R: uint8((img.w*img.h + x*y) % 255),
		G: uint8((img.w*img.h - x*y) % 255),
		B: uint8((w1 / h1 * x1 / y1) % 255),
		A: uint8((w1 * h1 / x1 * y1) % 255),
	}
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.w, img.h)
}

func (img Image) ColorModel() color.Model {
	return img.clrModel
}

func main() {
	m := Image{
		w:        300,
		h:        300,
		clrModel: color.RGBAModel,
	}

	pic.ShowImage(m)
}
