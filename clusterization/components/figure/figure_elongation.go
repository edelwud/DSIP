package figure

import "math"

type Elongation interface {
	FindElongation() int
	FindDiscreteCentralMoment(int, int) float64
}

// FindDiscreteCentralMoment finds discrete central moment for figure
func (f Figure) FindDiscreteCentralMoment(i int, j int) float64 {
	moment := 0.0

	cx, cy := f.FindCenterOfMass()
	for _, point := range f.Relative {
		moment += math.Pow(float64(point.X-cx), float64(i)) * math.Pow(float64(point.Y-cy), float64(j))
	}

	return moment
}

// FindElongation calculates elongation for figure
func (f Figure) FindElongation() int {
	m20 := f.FindDiscreteCentralMoment(2, 0)
	m02 := f.FindDiscreteCentralMoment(0, 2)
	m11 := f.FindDiscreteCentralMoment(1, 1)

	duplicate := math.Sqrt(math.Pow(m20-m02, 2) + 4*math.Pow(m11, 2))
	elongation := (m20 + m02 + duplicate) / (m20 - m02 - duplicate)

	return int(elongation)
}
