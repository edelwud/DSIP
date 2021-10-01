package figure

import (
	"image"
	"image/color"
	"math"
)

// Figure analysis methods implementation
type Figure struct {
	Snapshot *image.Gray
	Route    []image.Point
	Relative []image.Point
}

// CalculateRelative calculates Figure.Relative route
func (f *Figure) CalculateRelative() (width int, height int) {
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

	f.Relative = make([]image.Point, 0)

	for _, point := range f.Route {
		f.Relative = append(f.Relative, image.Point{
			X: point.X - int(minX),
			Y: point.Y - int(minY),
		})
	}

	return maxX, maxY
}

// DrawRoute draws Figure.Snapshot from route
func (f *Figure) DrawRoute() {
	width, height := f.CalculateRelative()
	f.Snapshot = image.NewGray(image.Rect(0, 0, width, height))

	for _, point := range f.Relative {
		f.Snapshot.Set(point.X, point.Y, color.Gray{Y: 255})
	}
}

// CreateFigure initializes Figure exemplar
func CreateFigure(route []image.Point) Figure {
	figure := Figure{
		Route:    route,
		Relative: make([]image.Point, 0),
	}
	figure.DrawRoute()
	return figure
}
