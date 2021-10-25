package hopfield

type Matrix struct {
	Data   [][]int
	Width  int
	Height int
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

func Transpose(matrix Matrix) Matrix {
	w := matrix.Width
	h := matrix.Height
	data := make([][]int, w)

	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			data[i] = append(data[i], matrix.Data[j][i])
		}
	}

	return Matrix{
		Data:   data,
		Width:  h,
		Height: w,
	}
}
