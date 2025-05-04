package core

import (
	"github.com/viterin/vek"
	"image"
	"image/color"
	"log"
	"math"
	"runtime"
	"sync"
	"time"
)

type Camera struct {
	Width, Height uint32
	Direction     []float64
}

func (c *Camera) Render(scene []RayMarchableObject) *image.RGBA {
	marcher := Marcher{
		MaximumNumberOfSteps: 100,
		Objects:              scene,
		DistanceLimit:        0.001,
	}

	lights := []Light{
		*NewLight(
			[]float64{
				-1.0, 0, 0.5,
			},
			color.RGBA{
				R: 255,
				G: 255,
				B: 255,
				A: 255,
			},
		),
	}

	result := image.NewRGBA(image.Rect(0, 0, int(c.Width), int(c.Height)))
	fov := math.Pi / 2
	aspect := float64(c.Width) / float64(c.Height)
	ch := make(chan int, runtime.NumCPU())
	var wg sync.WaitGroup
	// Generate camera rays
	for x := float64(0); x < float64(c.Width); x++ {
		ch <- 0
		wg.Add(1)
		go func() {
			start := time.Now()
			defer func() {
				log.Println("Rendering row took ", time.Since(start))
			}()
			for y := float64(0); y < float64(c.Height); y++ {
				px := ((x+0.5)/float64(c.Width))*2 - 1
				py := ((y+0.5)/float64(c.Height))*2 - 1
				pixelVec := []float64{
					px * aspect * math.Tan(fov/2),
					-py * math.Tan(fov/2),
					-1,
				}
				pixelVec = vek.DivNumber(pixelVec, vek.Norm(pixelVec))
				// Convert x and y to
				pixelRay := Ray{
					origin: []float64{
						0,
						0,
						0,
					},
					direction: pixelVec,
				}
				hit := marcher.March(pixelRay)
				if hit != nil {
					pixelColor := color.RGBA{
						A: 255,
					}
					hitNorm := CalculateNormal(hit.HitPosition, marcher.MinimumDistanceToPoint)
					for _, light := range lights {
						marcherWithLight := Marcher{
							MaximumNumberOfSteps: marcher.MaximumNumberOfSteps,
							Objects:              append(marcher.Objects, &light),
							DistanceLimit:        marcher.DistanceLimit,
						}
						dirVector := vek.Sub(light.Position, hit.HitPosition)
						dirVector = vek.DivNumber(dirVector, vek.Norm(dirVector))
						lightHit := marcherWithLight.March(Ray{
							origin:    vek.Add(hit.HitPosition, vek.MulNumber(dirVector, 0.1)),
							direction: dirVector,
						})
						if lightHit == nil || (*lightHit.HitObject).GetID() != light.GetID() {
							continue
						}
						intensity := math.Max(0, vek.Dot(hitNorm, dirVector))
						pixelColor = MixColors(pixelColor, color.RGBA{
							R: uint8(float64(light.Color.R) * intensity),
							G: uint8(float64(light.Color.G) * intensity),
							B: uint8(float64(light.Color.B) * intensity),
							A: 255,
						})
					}
					result.Set(int(x), int(y),
						pixelColor)
				} else {
					result.Set(int(x), int(y), color.Black)
				}
			}
			<-ch
			wg.Done()
		}()
	}
	wg.Wait()
	return result
}
