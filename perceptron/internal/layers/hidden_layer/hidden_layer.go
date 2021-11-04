package hidden_layer

import (
	"gonum.org/v1/gonum/mat"
	"perceptron/internal/layers"
)

type HiddenLayer struct {
	N *mat.VecDense
}

func (v HiddenLayer) Neurons() *mat.VecDense {
	return v.N
}

func (v HiddenLayer) CalculateWeights(layer layers.Layer) *mat.Dense {
	return nil
}

func NewHiddenLayer(vec *mat.VecDense) layers.Layer {
	return &HiddenLayer{N: vec}
}
