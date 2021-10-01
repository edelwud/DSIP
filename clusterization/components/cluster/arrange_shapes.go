package cluster

import (
	"clusterization/components/figure"
	"clusterization/components/utils"
	"image"
)

func ArrangeShapes(figures []figure.Figure, width int, height int) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for _, fig := range figures {
		colour := utils.GenerateRandomColour()
		for _, point := range fig.Route {
			img.Set(point.X, point.Y, colour)
		}
	}

	return img
}
