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

func (v *DistributionLayer) GenerateWeights(rows int, columns int) {
	matrix := make([]float64, rows*columns)
	for i := range matrix {
		x := 1.0
		if rand.Float64() > 0.5 {
			x = -1
		}
		matrix[i] = rand.Float64() * x
	}

	v.Weights = mat.NewDense(rows, columns, matrix)
}

func (v *DistributionLayer) GenerateThreshold() {
	thresholds := make([]float64, v.Neurons.Len())
	for i := 0; i < v.Neurons.Len(); i++ {
		x := 1.0
		if rand.Float64() > 0.5 {
			x = -1
		}
		thresholds[i] = rand.Float64() * x
	}

	v.Thresholds = mat.NewVecDense(len(thresholds), thresholds)
}

func (v DistributionLayer) W() *mat.Dense {
	return v.Weights
}

func (v DistributionLayer) N() *mat.VecDense {
	return v.Neurons
}

func (v DistributionLayer) Fill(vec *mat.VecDense) {
	v.Neurons.SetRawVector(vec.RawVector())
}

func (v DistributionLayer) Threshold() *mat.VecDense {
	return v.Thresholds
}

func NewDistributionLayer(length int) layers.Layer {
	vec := make([]float64, length)
	return &DistributionLayer{
		Neurons: mat.NewVecDense(length, vec),
	}
}
