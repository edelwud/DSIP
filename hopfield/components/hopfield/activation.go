package hopfield

func SignFunction(value int) (result int) {
	if value >= 0 {
		result = 1
	} else {
		result = -1
	}
	return
}

func Combinations(images ...Matrix) [][]Matrix {
	combinations := make([][]Matrix, 0)

	for i := 0; i < len(images); i++ {
		for j := i + 1; j < len(images); j++ {
			combinations = append(combinations, []Matrix{
				images[i],
				images[j],
			})
		}
	}

	return combinations
}

func GetWeights(images ...Matrix) Matrix {
	w := images[0].Width
	h := images[0].Height

	result := make([][]int, w)
	for i := range result {
		result[i] = make([]int, h)
	}

	for _, combination := range Combinations(images...) {
		m := Multiplication(combination[0], combination[1])
		result = Sum(m, Matrix{
			Data:   result,
			Width:  w,
			Height: h,
		}).Data
	}

	return Matrix{
		Data:   result,
		Width:  w,
		Height: h,
	}
}
