package hopfield

import (
	"gonum.org/v1/gonum/mat"
)

func GetWeights(images ...*mat.VecDense) *mat.Dense {
	l := images[0].Len()
	result := mat.NewDense(l, l, make([]float64, l*l))

	for i := range images {
		mul := mat.NewDense(l, l, make([]float64, l*l))
		mul.Mul(images[i], images[i].T())
		result.Add(result, mul)
	}

	for i := 0; i < l; i++ {
		result.Set(i, i, 0)
	}

	return result
}
