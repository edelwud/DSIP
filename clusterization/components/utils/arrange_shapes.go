package utils

import (
	"clusterization/components/figure"
	"image"
)

func ArrangeShapes(figures []figure.Figure, width int, height int) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for _, fig := range figures {
		colour := GenerateRandomColour()
		for _, point := range fig.Route {
			img.Set(point.X, point.Y, colour)
		}
	}

	return img
}
