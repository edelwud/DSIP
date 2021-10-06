package figure

import (
	"image"
	"image/color"
)

// Figure analysis methods implementation
type Figure struct {
	Image     image.Image
	Grayscale *image.Gray
	Snapshot  *image.Gray
	Route     []image.Point
	Relative  []image.Point
}

// DrawRoute draws Figure.Snapshot from route
func (f *Figure) DrawRoute() {
	min, max := f.FindDimensions()
	width, height := max.X-min.X, max.Y-min.Y

	relative := f.CalculateRelative()

	f.Snapshot = image.NewGray(image.Rect(0, 0, width, height))

	for _, point := range relative {
		f.Snapshot.Set(point.X, point.Y, color.Gray{Y: 255})
	}
}

// CreateFigure initializes Figure exemplar
func CreateFigure(img image.Image, grayscale *image.Gray, route []image.Point) Figure {
	figure := Figure{
		Image:     img,
		Grayscale: grayscale,
		Route:     route,
		Relative:  make([]image.Point, 0),
	}

	figure.DrawRoute()

	return figure
}
