package cluster

import (
	"clusterization/components/figure"
	"clusterization/components/mask"
	"clusterization/components/utils"
	"image"
	"image/color"
)

type Cluster interface {
	CollectAnalyticCriteria(criteria string) [][]*figure.Figure
	FindClusters([][]*figure.Figure) [][]*figure.Figure
	DrawAnalyticArea([][]*figure.Figure) image.Image
	DrawClusters(clusters [][]*figure.Figure, width int, height int) image.Image
}

type BasicCluster struct {
	Figures []figure.Figure
	Scale   int
}

const (
	CompactnessCriteria = "COMPACTNESS"
)

var analyticCriteria = map[string]func(f figure.Figure) (int, int){
	CompactnessCriteria: func(f figure.Figure) (int, int) {
		return f.FindPerimeter(), f.FindSquare()
	},
}

func (c BasicCluster) CollectAnalyticCriteria(criteria string) [][]*figure.Figure {
	maxX := 0
	maxY := 0

	for _, f := range c.Figures {
		x, y := analyticCriteria[criteria](f)
		if x > maxX {
			maxX = x
		}
		if y > maxY {
			maxY = y
		}
	}

	result := make([][]*figure.Figure, maxX+1)
	for i := range result {
		result[i] = make([]*figure.Figure, maxY+1)
	}

	for i := range c.Figures {
		x, y := analyticCriteria[criteria](c.Figures[i])
		result[x][y] = &c.Figures[i]
	}

	return result
}

func (c BasicCluster) DrawAnalyticArea(figures [][]*figure.Figure) image.Image {
	width, height := len(figures), len(figures[len(figures)-1])
	area := image.NewRGBA(image.Rect(0, 0, width+1, height+1))
	for x, row := range figures {
		for y, f := range row {
			if f != nil {
				area.Set(x, y, color.White)
			}
		}
	}
	return area
}

func (c BasicCluster) FindClusters(figures [][]*figure.Figure) [][]*figure.Figure {
	route := make([]image.Point, 0)
	result := make([][]*figure.Figure, 0)
	contourMask := mask.CreateContourAreaMask()
	m := contourMask.Generate(c.Scale)
	i := 0

	for {
		point := c.FindActivePoint(figures)
		if point == nil {
			break
		}
		x := point.X
		y := point.Y

		result = append(result, nil)

	loop:
		for {
			for _, diff := range m {
				x = point.X + diff.X
				y = point.Y + diff.Y

				if x < 0 {
					continue
				}

				if y < 0 {
					continue
				}

				if x >= len(figures) {
					continue
				}

				if y >= len(figures[x]) {
					continue
				}

				if figures[x][y] != nil {
					route = append(route, image.Point{
						X: x,
						Y: y,
					})
					result[i] = append(result[i], figures[x][y])
					figures[x][y] = nil
					continue loop
				}
			}
			break
		}
		i++
	}

	return result
}

func (c BasicCluster) DrawClusters(clusters [][]*figure.Figure, width int, height int) image.Image {
	result := image.NewRGBA(image.Rect(0, 0, width, height))

	for _, cluster := range clusters {
		c := utils.GenerateRandomColour()
		for _, f := range cluster {
			for _, point := range f.Route {
				result.Set(point.X, point.Y, c)
			}
		}
	}

	return result
}

func (c BasicCluster) FindActivePoint(figures [][]*figure.Figure) *image.Point {
	for x, row := range figures {
		for y := range row {
			if figures[x][y] != nil {
				return &image.Point{X: x, Y: y}
			}
		}
	}
	return nil
}

func CreateCluster(figures []figure.Figure, scale int) Cluster {
	return BasicCluster{Figures: figures, Scale: scale}
}
