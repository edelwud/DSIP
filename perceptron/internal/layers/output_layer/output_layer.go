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

func (v OutputLayer) GenerateWeights(rows int, columns int) {
	matrix := make([]float64, rows*columns)
	for i := range matrix {
		matrix[i] = rand.Float64()
	}

	v.Weights = mat.NewDense(rows, columns, matrix)
}

func (v OutputLayer) GenerateThreshold() {
	thresholds := make([]float64, v.Neurons.Len())
	for i := 0; i < v.Neurons.Len(); i++ {
		thresholds[i] = rand.Float64()
	}

	v.Thresholds = mat.NewVecDense(len(thresholds), thresholds)
}

func (v OutputLayer) Fill(vec *mat.VecDense) {
	v.Neurons.SetRawVector(vec.RawVector())
}

func NewOutputLayer(length int) layers.Layer {
	vec := make([]float64, length)
	return &OutputLayer{
		Neurons: mat.NewVecDense(length, vec),
	}
}
