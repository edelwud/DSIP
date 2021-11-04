package output_layer

import (
	"gonum.org/v1/gonum/mat"
	"perceptron/internal/layers"
)

type OutputLayer struct {
	N *mat.VecDense
}

func (v OutputLayer) Neurons() *mat.VecDense {
	return v.N
}

func (v OutputLayer) CalculateWeights(layer layers.Layer) *mat.Dense {
	return nil
}

func NewOutputLayer(vec *mat.VecDense) layers.Layer {
	return &OutputLayer{N: vec}
}
