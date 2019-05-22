package transform

import (
	"image"
)

type Rotate struct {}

func (r Rotate) Transform(i image.Image) image.Image {
	newBounds := image.Rectangle{i.Bounds().Min, image.Point{i.Bounds().Max.Y, i.Bounds().Max.X}}
	rotated := image.NewRGBA(newBounds)
	for y := i.Bounds().Min.Y; y < i.Bounds().Max.Y; y++ {
		for x := i.Bounds().Min.X; x < i.Bounds().Max.X; x++ {
			rotated.Set(y, x, i.At(x, newBounds.Max.X - y))
		}
	}
	return rotated.SubImage(newBounds)
}

func init() {
	r := new(Rotate)
	RegisterTransformer("rotate", r)
}