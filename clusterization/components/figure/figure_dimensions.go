package figure

import (
	"image"
	"math"
)

// Dimensions declares minimum interface for dimensions functionality implementation
type Dimensions interface {
	FindCenterOfMass() (int, int)
	FindDiscreteCentralMoment(int, int) int
	FindDimensions() (image.Point, image.Point)
	CalculateRelative() []image.Point
}

// FindDiscreteCentralMoment finds discrete central moment for figure
func (f Figure) FindDiscreteCentralMoment(i int, j int) int {
	moment := 0.0

	cx, cy := f.FindCenterOfMass()

	for _, point := range f.Relative {
		moment += math.Pow(float64(point.X-cx), float64(i)) * math.Pow(float64(point.Y-cy), float64(j))
	}

	return int(moment)
}

// FindCenterOfMass finds center of mass for figure
func (f Figure) FindCenterOfMass() (xSummary int, ySummary int) {
	for _, point := range f.Relative {
		x, y := point.X, point.Y
		xSummary += x
		ySummary += y
	}

	xSummary /= len(f.Relative)
	ySummary /= len(f.Relative)
	return
}

// FindDimensions finds figure dimensions
func (f Figure) FindDimensions() (min image.Point, max image.Point) {
	max.X = 0
	max.Y = 0

	min.X = (1 << 31) - 1
	min.Y = (1 << 31) - 1

	for _, point := range f.Route {
		x, y := point.X, point.Y
		if x > max.X {
			max.X = x
		}
		if y > max.Y {
			max.Y = y
		}
		if x < min.X {
			min.X = x
		}
		if y < min.Y {
			min.Y = y
		}
	}

	return
}

// CalculateRelative calculates Figure.Relative route
func (f *Figure) CalculateRelative() []image.Point {
	min, _ := f.FindDimensions()

	f.Relative = make([]image.Point, 0)

	for _, point := range f.Route {
		f.Relative = append(f.Relative, image.Point{
			X: point.X - min.X,
			Y: point.Y - min.Y,
		})
	}

	return f.Relative
}
