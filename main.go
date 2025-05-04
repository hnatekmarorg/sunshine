package main

import (
	"github.com/hnatekmarorg/sunshine/core"
	"image/png"
	"os"
)

func main() {
	c := core.Camera{
		Width:     512,
		Height:    512,
		Direction: []float64{0, 0, 1},
	}
	image := c.Render([]core.RayMarchableObject{
		core.NewSphere([]float64{
			0, 0, -2,
		}, 1),
		core.NewSphere([]float64{
			1, 0, -2,
		}, 1),
		core.NewSphere([]float64{
			0.5, 1, -2,
		}, 1),
		core.NewSphere([]float64{
			0, 0, -0.5,
		}, 0.1),
	})
	f, _ := os.Create("output.png")
	defer f.Close()
	png.Encode(f, image)
}
