package cluster

import (
	"clusterization/components/figure"
	"image"
	"image/color"
)

type ConnectedAreas struct {
	Image *image.Gray
}

func (a ConnectedAreas) FindConnectedAreas() []figure.Figure {
	figures := make([]figure.Figure, 0)
	//for a.HasArea() {
	route := a.WalkThroughArea()
	a.ClearRoute(route)
	//}

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

	for {
		tl := a.Image.GrayAt(x, y).Y
		tr := a.Image.GrayAt(x+1, y).Y
		trr := a.Image.GrayAt(x+2, y).Y
		br := a.Image.GrayAt(x+1, y+1).Y
		brr := a.Image.GrayAt(x+2, y+1).Y
		bbr := a.Image.GrayAt(x+1, y+2).Y
		bbrr := a.Image.GrayAt(x+2, y+2).Y
		bl := a.Image.GrayAt(x, y+1).Y
		bbl := a.Image.GrayAt(x, y+2).Y

		if tl == 255 {
			a.Image.Set(x, y, color.Gray{Y: 0})
			route = append(route, image.Point{
				X: x,
				Y: y,
			})
		}

		if tr == 255 {
			x++
			continue
		}

		if br == 255 {
			x++
			y++
			continue
		}

		if bl == 255 {
			y++
			continue
		}

		if trr == 255 {
			x += 2
			continue
		}

		if brr == 255 {
			x += 2
			y += 1
			continue
		}

		if bbrr == 255 {
			x += 2
			y += 2
			continue
		}

		if bbr == 255 {
			x += 1
			y += 2
			continue
		}

		if bbl == 255 {
			y += 2
			continue
		}

		break
	}

	return route
}

func (a ConnectedAreas) GenerateMask(x int, y int) [][]int {
	result := make([][]int, 0)

	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			result[i] = append(result[i], j)
		}
		result = append(result, []int{})
	}

	return result
}

func (a ConnectedAreas) ClearRoute(route []image.Point) {
	for _, point := range route {
		x, y := point.X, point.Y
		a.Image.Set(x, y, color.Gray{Y: 0})
	}
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
	return ConnectedAreas{Image: img}
}
