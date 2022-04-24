/*

Exercise: Images

Remember the picture generator you wrote earlier? Let's write another one, but this time it will return an implementation of image.Image instead of a slice of data.

Define your own Image type, implement the necessary methods, and call pic.ShowImage.

Bounds should return a image.Rectangle, like image.Rect(0, 0, w, h).

ColorModel should return color.RGBAModel.

At should return a color; the value v in the last picture generator corresponds to color.RGBA{v, v, 255, 255} in this one.

*/

package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	dx, dy int
	img    [][]uint8
}

func NewImage(dx, dy int) Image {
	t := Image{dx, dy, nil}
	t.img = make([][]uint8, dy)
	for y := range t.img {
		t.img[y] = make([]uint8, dx)
		for x := range t.img[y] {
			t.img[y][x] = uint8((x + y) / 2)
		}
	}
	return t
}

func (t Image) ColorModel() color.Model {
	return color.GrayModel
}

func (t Image) Bounds() image.Rectangle {
	return image.Rectangle{image.Point{0, 0}, image.Point{t.dx, t.dy}}
}

func (t Image) At(x, y int) color.Color {
	return color.Gray{t.img[y][x]}
}

func main() {
	m := NewImage(400, 300)
	pic.ShowImage(m)
}
