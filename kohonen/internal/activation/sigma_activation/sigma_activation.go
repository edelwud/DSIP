package sigma_activation

import (
	"gonum.org/v1/gonum/mat"
	"kohonen/internal/activation"
	"math"
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
		result = append(result, x*(1-x))
	}

	return mat.NewVecDense(len(result), result)
}

func (a SigmaActivation) DerivativeValue(x float64) float64 {
	return x * (1 - x)
}

func NewSigmaActivation() activation.Activation {
	return &SigmaActivation{}
}
