package figure

import (
	"image"
	"math"
)

type Figure struct {
	Snapshot *image.Gray
	Route    []image.Point
	Relative []image.Point
}

func (f Figure) CalculateRelative() {
	maxX := 0
	maxY := 0
	minX := math.Inf(1)
	minY := math.Inf(1)

	for _, point := range f.Route {
		x, y := point.X, point.Y
		if x > maxX {
			maxX = x
		}
		if y > maxY {
			maxY = y
		}
		if float64(x) < minX {
			minX = float64(x)
		}
		if float64(y) < minY {
			minY = float64(y)
		}
	}

	for _, point := range f.Route {
		f.Relative = append(f.Relative, image.Point{
			X: point.X - int(minX),
			Y: point.Y - int(minY),
		})
	}
}

func CreateFigure(snapshot *image.Gray, route []image.Point) Figure {
	return Figure{
		Snapshot: snapshot,
		Route:    route,
		Relative: make([]image.Point, 0),
	}
}
