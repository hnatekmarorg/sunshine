package core

import "github.com/viterin/vek"

var nextObjectID uint64

type RayMarchableObject interface {
	// SDF Signed distance function
	SDF(position []float64) float64
	GetID() uint64
}

func CalculateNormal(hitPoint []float64, f func(point []float64) float64) []float64 {
	const epsilon = 0.0001
	diff := []float64{
		f(vek.Add(hitPoint, []float64{epsilon, 0, 0})) - f(vek.Sub(hitPoint, []float64{epsilon, 0, 0})),
		f(vek.Add(hitPoint, []float64{0, epsilon, 0})) - f(vek.Sub(hitPoint, []float64{0, epsilon, 0})),
		f(vek.Add(hitPoint, []float64{0, 0, epsilon})) - f(vek.Sub(hitPoint, []float64{0, 0, epsilon})),
	}
	diff = vek.DivNumber(diff, vek.Norm(diff))
	return diff
}
