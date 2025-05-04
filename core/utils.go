package core

import (
	"image/color"
	"math"
)

func MixColors(a color.RGBA, b color.RGBA) color.RGBA {
	return color.RGBA{
		R: uint8(math.Min(float64(a.R+b.R), 255)),
		G: uint8(math.Min(float64(a.G+b.G), 255)),
		B: uint8(math.Min(float64(a.B+b.B), 255)),
		A: uint8(math.Min(float64(a.A+b.A), 255)),
	}
}
