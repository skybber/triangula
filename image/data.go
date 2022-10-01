package image

import "github.com/skybber/Triangula/color"

type Data interface {
	RGBAt(x, y int) color.RGB
	Size() (int, int)
}
