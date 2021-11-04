package layers

import "gonum.org/v1/gonum/mat"

type Layer interface {
	Neurons() *mat.VecDense
	CalculateWeights(layer Layer) *mat.Dense
}
