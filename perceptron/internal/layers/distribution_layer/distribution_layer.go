package distribution_layer

import (
	"gonum.org/v1/gonum/mat"
	"math/rand"
	"perceptron/internal/layers"
)

type DistributionLayer struct {
	Neurons    *mat.VecDense
	Thresholds *mat.VecDense
	Weights    *mat.Dense
}

func (v DistributionLayer) GenerateWeights(rows int, columns int) {
	matrix := make([]float64, rows*columns)
	for i := range matrix {
		matrix[i] = rand.Float64()
	}

	v.Weights = mat.NewDense(rows, columns, matrix)
}

func (v DistributionLayer) GenerateThreshold() {
	thresholds := make([]float64, v.Neurons.Len())
	for i := 0; i < v.Neurons.Len(); i++ {
		thresholds[i] = rand.Float64()
	}

	v.Thresholds = mat.NewVecDense(len(thresholds), thresholds)
}

func NewDistributionLayer(length int) layers.Layer {
	vec := make([]float64, length)
	return &DistributionLayer{
		Neurons: mat.NewVecDense(length, vec),
	}
}
