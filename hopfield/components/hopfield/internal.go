package hopfield

import (
	mat "github.com/gonum/matrix/mat64"
	"image"
	"image/color"
)

const (
	PositiveNormalization = 1.0
	NegativeNormalization = -1.0

	PositiveAlias = 255
)

func NormalizeObject(img *image.Gray) []float64 {
	x, y := img.Bounds().Dx(), img.Bounds().Dy()
	result := make([]float64, 0)

	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			c := img.GrayAt(j, i).Y
			normalized := NegativeNormalization
			if c == PositiveAlias {
				normalized = PositiveNormalization
			}
			result = append(result, normalized)
		}
	}

	return result
}

func DenormalizeObject(m *mat.Dense) image.Image {
	r, c := m.Dims()
	symbol := image.NewGray(image.Rect(0, 0, c, r))

	x, y := c, r

	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			value := m.At(j, i)
			c := uint8(255)
			if value == -1 {
				c = 0
			}

			symbol.Set(i, j, color.Gray{
				Y: c,
			})
		}
	}

	return symbol
}
