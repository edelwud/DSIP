package cluster

import (
	"clusterization/components/figure"
	"clusterization/components/mask"
	"image"
	"image/color"
)

// ConnectedAreas finds connected areas in binarized images
type ConnectedAreas struct {
	Image   *image.Gray
	Figures []figure.Figure
	Mask    mask.AreaMask
	Scale   int
}

// FindConnectedAreas processes image scanning to find connected areas
func (a *ConnectedAreas) FindConnectedAreas() []figure.Figure {
	for a.HasArea() {
		route := a.WalkThroughArea()
		a.Figures = append(a.Figures, figure.CreateFigure(route))
	}

	return a.Figures
}

// GetStartPoint returns start point for WalkThroughArea
func (a ConnectedAreas) GetStartPoint() *image.Point {
	for y := a.Image.Bounds().Min.Y; y < a.Image.Bounds().Max.Y; y++ {
		for x := a.Image.Bounds().Min.X; x < a.Image.Bounds().Max.X; x++ {
			if a.Image.GrayAt(x, y).Y == 255 {
				return &image.Point{
					X: x,
					Y: y,
				}
			}
		}
	}
	return nil
}

// WalkThroughArea walks through connected area, returns path
func (a ConnectedAreas) WalkThroughArea() []image.Point {
	route := make([]image.Point, 0)

	point := a.GetStartPoint()
	if point == nil {
		return nil
	}

	x, y := point.X, point.Y

	m := a.Mask.Generate(a.Scale)

loop:
	for {
		for _, diff := range m {
			intensity := a.Image.GrayAt(x+diff.X, y+diff.Y).Y
			if intensity == 255 {
				x += diff.X
				y += diff.Y
				a.Image.Set(x, y, color.Gray{Y: 0})
				route = append(route, image.Point{
					X: x,
					Y: y,
				})
				continue loop
			}
		}

		for _, r := range route {
			for _, diff := range m {
				intensity := a.Image.GrayAt(r.X+diff.X, r.Y+diff.Y).Y
				if intensity == 255 {
					x = r.X + diff.X
					y = r.Y + diff.Y
					continue loop
				}
			}
		}

		a.Image.Set(x, y, color.Gray{Y: 200})
		break
	}

	return route
}

// ClearRoute clears route in parent image
func (a ConnectedAreas) ClearRoute(route []image.Point) {
	for _, point := range route {
		x, y := point.X, point.Y
		a.Image.Set(x, y, color.Gray{Y: 0})
	}
}

// HasArea looks for white areas in parent image
func (a ConnectedAreas) HasArea() bool {
	for y := a.Image.Bounds().Min.Y; y < a.Image.Bounds().Max.Y; y++ {
		for x := a.Image.Bounds().Min.X; x < a.Image.Bounds().Max.X; x++ {
			intensity := a.Image.GrayAt(x, y).Y
			if intensity == 255 {
				return true
			}
		}
	}
	return false
}

// CreateConnectedAreasAnalyzer creates ConnectedAreas exemplar
func CreateConnectedAreasAnalyzer(img *image.Gray, mask mask.AreaMask, scale int) ConnectedAreas {
	return ConnectedAreas{
		Image:   img,
		Mask:    mask,
		Scale:   scale,
		Figures: make([]figure.Figure, 0),
	}
}
