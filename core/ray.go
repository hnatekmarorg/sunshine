package core

import "github.com/viterin/vek"

type Ray struct {
	origin, direction []float64
}

// t returns point based on t param
func (r *Ray) t(t float64) []float64 {
	return vek.Add(r.origin, vek.MulNumber(r.direction, t))
}

func (r *Ray) Init(origin []float64, direction []float64) Ray {
	if len(origin) != 3 {
		panic("Invalid origin length must be 3 dimensional vector")
	}
	if len(direction) != 3 {
		panic("Invalid direction length must be 3 dimensional vector")
	}
	return Ray{
		origin:    origin,
		direction: direction,
	}
}
