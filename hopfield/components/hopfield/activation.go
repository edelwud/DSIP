package hopfield

import (
	"gonum.org/v1/gonum/mat"
)

func SignFunction(value float64) (result float64) {
	if value >= 0 {
		result = 1
	} else {
		result = -1
	}
	return
}

func Activation(m *mat.VecDense) *mat.VecDense {
	result := mat.NewVecDense(m.Len(), make([]float64, m.Len()))

	for i := 0; i < m.Len(); i++ {
		result.SetVec(i, SignFunction(m.AtVec(i)))
	}

	return result
}

func CompareWithGolden(img *mat.VecDense, golden []*mat.VecDense) bool {
	for _, gold := range golden {
		if Compare(gold, img) {
			return true
		}
	}
	return false
}

func Compare(m1 *mat.VecDense, m2 *mat.VecDense) bool {
	r, c := m1.Dims()

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if m1.At(i, j) != m2.At(i, j) {
				return false
			}
		}
	}

	return true
}
