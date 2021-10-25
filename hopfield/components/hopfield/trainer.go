package hopfield

import (
	mat "github.com/gonum/matrix/mat64"
)

func GetWeights(images ...*mat.Dense) *mat.Dense {
	r, _ := images[0].T().Dims()
	_, c := images[1].Dims()
	result := mat.NewDense(r, c, make([]float64, r*c))

	for _, golden := range images {
		mul := mat.NewDense(r, c, make([]float64, r*c))
		mul.Mul(golden.T(), golden)
		result.Add(result, mul)
	}

	for i := 0; i < r; i++ {
		result.Set(i, i, 0)
	}

	return result
}
