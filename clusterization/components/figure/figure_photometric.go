package figure

import (
	"image/color"
)

type Photometric interface {
	AverageIntensity() uint8
	AverageColor() color.RGBA
	IntensityHistogram() []int
	IntensityDispersion() []int
}

func (f Figure) AverageIntensity() uint8 {
	intensity := 0.0
	size := float64(len(f.Route))

	for _, point := range f.Route {
		intensity += float64(f.Grayscale.GrayAt(point.X, point.Y).Y)
	}

	return uint8(intensity / size)
}

func (f Figure) AverageColor() color.RGBA {
	r := 0.0
	g := 0.0
	b := 0.0
	aa := 0.0

	for _, point := range f.Route {
		red, green, blue, a := f.Image.At(point.X, point.Y).RGBA()
		red, green, blue, a = red>>8, green>>8, blue>>8, a>>8

		r += float64(red)
		g += float64(green)
		b += float64(blue)
		aa += float64(a)
	}

	return color.RGBA{
		R: uint8(r),
		G: uint8(g),
		B: uint8(b),
		A: uint8(aa),
	}
}

func (f Figure) IntensityHistogram() []int {
	histogram := make([]int, 256)

	for _, point := range f.Route {
		intensity := f.Grayscale.GrayAt(point.X, point.Y).Y
		histogram[intensity]++
	}

	return histogram
}

func (f Figure) IntensityDispersion() []int {
	dispersion := make([]int, 0)
	averageIntensity := f.AverageIntensity()

	for _, point := range f.Route {
		intensity := f.Grayscale.GrayAt(point.X, point.Y).Y
		dispersion = append(dispersion, int(intensity-averageIntensity))
	}

	return dispersion
}
