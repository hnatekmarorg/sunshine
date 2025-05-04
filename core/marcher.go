package core

import (
	"github.com/viterin/vek"
	"image/color"
	"math"
	"sync/atomic"
)

type Light struct {
	RayMarchableObject
	Color    color.RGBA
	Position []float64
	id       uint64
}

func NewLight(position []float64, col color.RGBA) *Light {
	return &Light{
		RayMarchableObject: nil,
		Color:              col,
		Position:           position,
		id:                 atomic.AddUint64(&nextObjectID, 1),
	}
}

func (l *Light) GetID() uint64 {
	return l.id
}

func (light *Light) SDF(point []float64) float64 {
	distanceToCenter := vek.Sub(point, light.Position)
	return math.Sqrt(math.Pow(distanceToCenter[0], 2) + math.Pow(distanceToCenter[1], 2) + math.Pow(distanceToCenter[2], 2))
}

type Marcher struct {
	MaximumNumberOfSteps uint
	Objects              []RayMarchableObject
	// RayMarcher will stop if SDF is smaller than this
	DistanceLimit float64
}

type MarchHit struct {
	HitPosition []float64
	Distance    float64
	HitObject   *RayMarchableObject
}

func (m *Marcher) MinimumDistanceToPoint(point []float64) float64 {
	minimumDistance := math.Inf(1)
	for _, object := range m.Objects {
		minimumDistance = math.Min(object.SDF(point), minimumDistance)
	}
	return minimumDistance
}

func (m *Marcher) March(ray Ray) *MarchHit {
	distance := 0.0
	var hitObject *RayMarchableObject = nil
	for i := uint(0); i < m.MaximumNumberOfSteps; i++ {
		minimumDistance := math.Inf(1)
		for oID, object := range m.Objects {
			distanceToObject := object.SDF(ray.origin)
			if minimumDistance > distanceToObject {
				minimumDistance = distanceToObject
				hitObject = &m.Objects[oID]
			}
		}
		ray = Ray{
			ray.t(minimumDistance),
			ray.direction,
		}
		distance += minimumDistance
		if minimumDistance < m.DistanceLimit {
			return &MarchHit{
				HitPosition: ray.origin,
				Distance:    distance,
				HitObject:   hitObject,
			}
		}
	}
	return nil
}
