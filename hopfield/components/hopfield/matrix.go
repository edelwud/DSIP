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

func Multiplication(m1 Matrix, m2 Matrix) Matrix {
	w := m1.Width
	h := m2.Height
	result := make([][]int, w)

	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			result[i] = append(result[i], 0)
			for k := 0; k < w; k++ {
				result[i][j] += m1.Data[i][k] * m2.Data[k][j]
			}
		}
	}

	return Matrix{
		Data:   result,
		Width:  w,
		Height: h,
	}
}
