package activation

import "gonum.org/v1/gonum/mat"

type Activation interface {
	Apply(vec *mat.VecDense) *mat.VecDense
	Derivative(vec *mat.VecDense) *mat.VecDense
}
