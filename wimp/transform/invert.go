package transform

import (
	"image"
	"image/color"
)

type Invert struct {}

func (r Invert) invert(c color.Color) color.Color {
	R, G, B, A := c.RGBA()
	return color.NRGBA{R: 255 - uint8(R), G: 255 - uint8(G), B: 255 - uint8(B), A: uint8(A)}
}

func (r Invert) Transform(i image.Image) image.Image {
	inverted := image.NewRGBA(i.Bounds())
	for y := i.Bounds().Min.Y; y < i.Bounds().Max.Y; y++ {
		for x := i.Bounds().Min.X; x < i.Bounds().Max.X; x++ {
			inverted.Set(x, y, r.invert(i.At(x, y)))
		}
	}
	return inverted.SubImage(i.Bounds())
}

func init() {
	i := new(Invert)
	RegisterTransformer("invert", i)
}