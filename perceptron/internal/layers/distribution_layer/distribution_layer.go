package distribution_layer

import (
	"gonum.org/v1/gonum/mat"
	"perceptron/internal/layers"
)

type DistributionLayer struct {
	N *mat.VecDense
}

func (v DistributionLayer) Neurons() *mat.VecDense {
	return v.N
}

func (v DistributionLayer) CalculateWeights(layer layers.Layer) *mat.Dense {
	return nil
}

func NewDistributionLayer(vec *mat.VecDense) layers.Layer {
	return &DistributionLayer{N: vec}
}
