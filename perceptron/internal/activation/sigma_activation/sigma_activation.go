package sigma_activation

import (
	"gonum.org/v1/gonum/mat"
	"math"
	"perceptron/internal/activation"
)

type SigmaActivation struct {
}

const Alpha = -1

func (a SigmaActivation) Apply(vec *mat.VecDense) *mat.VecDense {
	result := make([]float64, 0)

	for i := 0; i < vec.Len(); i++ {
		x := vec.AtVec(i)
		exp := math.Pow(math.E, Alpha*x)
		result = append(result, 1/(1+exp))
	}

	return mat.NewVecDense(len(result), result)
}

func (a SigmaActivation) ApplyValue(x float64) float64 {
	return 1 / (1 + math.Pow(math.E, Alpha*x))
}

func (a SigmaActivation) Derivative(vec *mat.VecDense) *mat.VecDense {
	result := make([]float64, 0)

	for i := 0; i < vec.Len(); i++ {
		x := vec.AtVec(i)
		if x < 0 {
			x *= 0.01
		} else if x > 1 {
			x = 1 + 0.01*(x-1)
		}
		result = append(result, x)
	}

	return mat.NewVecDense(len(result), result)
}

func (a SigmaActivation) DerivativeValue(x float64) float64 {
	if x < 0 {
		x *= 0.01
	} else {
		x = 1 + 0.01*(x-1)
	}
	return x
}

func NewSigmaActivation() activation.Activation {
	return &SigmaActivation{}
}
