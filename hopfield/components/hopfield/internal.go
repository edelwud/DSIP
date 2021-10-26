package hopfield

import (
	"gonum.org/v1/gonum/mat"
	"image"
	"image/color"
)

const (
	PositiveNormalization = 1.0
	NegativeNormalization = -1.0

	PositiveAlias = 0
	NegativeAlias = 255
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

func DenormalizeObject(width int, height int, m *mat.VecDense) *image.Gray {
	symbol := image.NewGray(image.Rect(0, 0, width, height))

	x, y := width, height

	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			value := m.AtVec(width*j + i)
			c := uint8(PositiveAlias)
			if value == NegativeNormalization {
				c = NegativeAlias
			}

			symbol.Set(i, j, color.Gray{
				Y: c,
			})
		}
	}

	return symbol
}

func ConvertToVec(img *image.Gray) *mat.VecDense {
	normalized := NormalizeObject(img)
	vec := mat.NewVecDense(len(normalized), normalized)
	return vec
}
