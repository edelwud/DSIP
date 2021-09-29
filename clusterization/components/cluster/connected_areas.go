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

loop:
	for {
		for k := 1; k < 3; k++ {
			for i := -k; i < k+1; i++ {
				for j := -k; j < k+1; j++ {
					intensity := a.Image.GrayAt(x+j, y+i).Y
					if intensity == 255 {
						x += j
						y += i
						a.Image.Set(x, y, color.Gray{Y: 0})
						route = append(route, image.Point{
							X: x,
							Y: y,
						})
						continue loop
					}
				}
			}
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
