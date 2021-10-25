package hopfield

import "image"

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
			color := img.GrayAt(j, i).Y
			normalized := NegativeNormalization
			if color == PositiveAlias {
				normalized = PositiveNormalization
			}
			result = append(result, normalized)
		}
	}

	return result
}
