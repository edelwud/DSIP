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

func (f Figure) FindSquare() int {
	return len(f.Relative)
}

func (f Figure) FindPerimeter() int {
	_, max := f.FindDimensions()
	return 2*max.X + 2*max.Y
}

func (f Figure) FindDimensions() (min image.Point, max image.Point) {
	max.X = 0
	max.Y = 0

	min.X = math.MaxInt
	min.Y = math.MaxInt

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

// DrawRoute draws Figure.Snapshot from route
func (f *Figure) DrawRoute() {
	min, max := f.FindDimensions()
	width, height := max.X-min.X, max.Y-min.Y

	relative := f.CalculateRelative()
	x, y := f.FindCenterOfMass()

	println(width, height)

	f.Snapshot = image.NewGray(image.Rect(0, 0, width, height))

	for _, point := range relative {
		f.Snapshot.Set(point.X, point.Y, color.Gray{Y: 100})
	}

	f.Snapshot.Set(x, y, color.Gray{Y: 255})
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
