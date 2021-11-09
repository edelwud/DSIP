package activation

import "gonum.org/v1/gonum/mat"

type Activation interface {
	Apply(vec *mat.VecDense) *mat.VecDense
	ApplyValue(x float64) float64
	Derivative(vec *mat.VecDense) *mat.VecDense
	DerivativeValue(x float64) float64
}
