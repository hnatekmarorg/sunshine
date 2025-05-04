package core

import (
	"github.com/viterin/vek"
	"math"
	"sync/atomic"
)

type Sphere struct {
	RayMarchableObject
	Position []float64
	Radius   float64
	id       uint64
}

func NewSphere(position []float64, radius float64) *Sphere {
	return &Sphere{
		Position: position,
		Radius:   radius,
		id:       atomic.AddUint64(&nextObjectID, 1),
	}
}

func (s *Sphere) GetID() uint64 {
	return s.id
}

func (s *Sphere) SDF(point []float64) float64 {
	distanceToCenter := vek.Sub(point, s.Position)
	return math.Sqrt(math.Pow(distanceToCenter[0], 2)+math.Pow(distanceToCenter[1], 2)+math.Pow(distanceToCenter[2], 2)) - s.Radius
}
