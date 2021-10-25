package hopfield

import "image"

type Matrix struct {
	Data   [][]int
	Width  int
	Height int
}

const (
	PositiveNormalization = 1
	NegativeNormalization = -1

	PositiveAlias = 255
)

func NormalizeObject(img *image.Gray) []int {
	x, y := img.Bounds().Dx(), img.Bounds().Dy()
	result := make([]int, 0)

	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			color := img.GrayAt(i, j).Y
			normalized := NegativeNormalization
			if color == PositiveAlias {
				normalized = PositiveNormalization
			}
			result = append(result, normalized)
		}
	}

	return result
}

func ConvertToMatrix(width int, height int, data []int) Matrix {
	matrix := make([][]int, width)

	for i := range matrix {
		for j := 0; j < height; j++ {
			offset := j*width + i
			matrix[i] = append(matrix[i], data[offset])
		}
	}

	return Matrix{
		Data:   matrix,
		Width:  width,
		Height: height,
	}
}

func ZeroingMainDiagonal(matrix Matrix) Matrix {
	x := matrix.Width

	i := 0
	j := 0

	for i < x {
		matrix.Data[i][j] = 0
		i++
		j++
	}

	return matrix
}
