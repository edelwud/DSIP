package cluster

import (
	"clusterization/components/figure"
	"fmt"
	"image"
	"image/color"
	"math"
)

type ConnectedAreas struct {
	Image *image.Gray
	Areas []*image.Gray
}

func (a *ConnectedAreas) FindConnectedAreas() []figure.Figure {
	figures := make([]figure.Figure, 0)

	for a.HasArea() {
		route := a.WalkThroughArea()
		a.ClearRoute(route)
		a.Areas = append(a.Areas, a.DrawRoute(route))
	}

	return figures
}

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

func (a ConnectedAreas) WalkThroughArea() []image.Point {
	route := make([]image.Point, 0)

	point := a.GetStartPoint()
	if point == nil {
		return nil
	}

	x, y := point.X, point.Y

	mask := a.GenerateMask(30)

loop:
	for {
		for _, diff := range mask {
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
		break
	}

	return route
}

func (a ConnectedAreas) GenerateMask(scale int) []image.Point {
	result := make([]image.Point, 0)

	result = append(result, image.Point{X: 0, Y: 0})

	x := 0
	y := 0

	width := 1
	height := 1
	pawn := true

	for width <= scale {
		xFlag := false
		yFlag := true
		flag := -width - (width - 1)
		x++

		result = append(result, image.Point{X: x, Y: 1 - width})

		for x != width || y != -height {
			if flag > 0 {
				if x == width {
					xFlag = false
				}
				if x == -width {
					xFlag = true
				}
				if xFlag {
					x++
				} else {
					x--
				}
				pawn = true
				flag--
			} else {
				if y == height {
					yFlag = false
				}
				if y == -height {
					yFlag = true
				}
				if yFlag {
					y++
				} else {
					y--
				}
				pawn = false
				flag++
			}
			result = append(result, image.Point{X: x, Y: y})
			if flag == 0 {
				if pawn {
					flag = -2 * width
				} else {
					flag = +2 * width
				}
			}
		}
		width++
		height++
	}

	return result
}

func (a ConnectedAreas) ClearRoute(route []image.Point) {
	for _, point := range route {
		x, y := point.X, point.Y
		a.Image.Set(x, y, color.Gray{Y: 0})
	}
}

func (a ConnectedAreas) DrawRoute(route []image.Point) *image.Gray {
	maxX := 0
	maxY := 0
	minX := math.Inf(1)
	minY := math.Inf(1)
	for _, point := range route {
		x, y := point.X, point.Y
		if x > maxX {
			maxX = x
		}
		if y > maxY {
			maxY = y
		}
		if float64(x) < minX {
			minX = float64(x)
			fmt.Println(x)
		}
		if float64(y) < minY {
			minY = float64(y)
		}
	}

	for i, point := range route {
		route[i] = image.Point{
			X: point.X - int(minX),
			Y: point.Y - int(minY),
		}
	}

	img := image.NewGray(image.Rect(0, 0, maxX-int(minX), maxY-int(minY)))

	for _, point := range route {
		x, y := point.X, point.Y
		img.Set(x, y, color.Gray{Y: 255})
	}

	return img
}

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

func CreateConnectedAreasAnalyzer(img *image.Gray) ConnectedAreas {
	return ConnectedAreas{Image: img, Areas: make([]*image.Gray, 0)}
}
