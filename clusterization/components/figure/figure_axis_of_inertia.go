package figure

import "math"

type AxisOfInertia interface {
	FindAxisOfInertia() float64
}

// FindAxisOfInertia finds orientation of the main axis of inertia
func (f Figure) FindAxisOfInertia() float64 {
	m11 := f.FindDiscreteCentralMoment(1, 1)
	m20 := f.FindDiscreteCentralMoment(2, 0)
	m02 := f.FindDiscreteCentralMoment(0, 2)

	return math.Atan((2*m11)/(m20-m02)) / 2
}
