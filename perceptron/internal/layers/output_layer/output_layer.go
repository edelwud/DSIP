package output_layer

import (
	"gonum.org/v1/gonum/mat"
	"math/rand"
	"perceptron/internal/layers"
)

type OutputLayer struct {
	Neurons    *mat.VecDense
	Thresholds *mat.VecDense
	Weights    *mat.Dense
}

func (v *OutputLayer) GenerateWeights(rows int, columns int) {
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

func (v *OutputLayer) GenerateThreshold() {
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

func (v OutputLayer) Fill(vec *mat.VecDense) {
	v.Neurons.SetRawVector(vec.RawVector())
}

func (v OutputLayer) W() *mat.Dense {
	return v.Weights
}

func (v OutputLayer) N() *mat.VecDense {
	return v.Neurons
}

func (v OutputLayer) Threshold() *mat.VecDense {
	return v.Thresholds
}

func NewOutputLayer(length int) layers.Layer {
	vec := make([]float64, length)
	return &OutputLayer{
		Neurons: mat.NewVecDense(length, vec),
	}
}
