package rbf

import (
	"gonum.org/v1/gonum/mat"
	"math"
	"math/rand"
)

type Layers struct {
	LayersNum int
	Sizes     []int

	Neurons    []*mat.VecDense
	LastChange *mat.Dense
	Weights    *mat.Dense
	Bias       *mat.VecDense
}

const (
	LayersNum = 2
)

func (l *Layers) InitLayers() {
	neurons := make([]*mat.VecDense, l.LayersNum)

	for i := range neurons {
		n := make([]float64, l.Sizes[i])
		neurons[i] = mat.NewVecDense(l.Sizes[i], n)
	}

	l.Neurons = neurons
}

func (l *Layers) InitWeights() {
	w := make([]float64, l.Sizes[1]*l.Sizes[0])
	for j := range w {
		sign := rand.Float64()
		if sign > 0.5 {
			w[j] = rand.Float64()
		} else {
			w[j] = -rand.Float64()
		}
	}
	l.Weights = mat.NewDense(l.Sizes[1], l.Sizes[0], w)
}

func (l *Layers) InitBias() {
	b := make([]float64, l.Sizes[1])
	for j := range b {
		sign := rand.Float64()
		if sign > 0.5 {
			b[j] = rand.Float64()
		} else {
			b[j] = -rand.Float64()
		}
	}

	l.Bias = mat.NewVecDense(l.Sizes[1], b)
}

func (l *Layers) InitLastChange() {
	e := make([]float64, l.Sizes[1]*l.Sizes[0])
	l.LastChange = mat.NewDense(l.Sizes[1], l.Sizes[0], e)
}

func (l Layers) FindResult() int {
	max := -math.MaxFloat64
	result := -1
	outputLayer := l.Neurons[l.LayersNum-1]

	for i := 0; i < outputLayer.Len(); i++ {
		x := outputLayer.AtVec(i)
		if x > max {
			max = x
			result = i
		}
	}

	return result
}

func (l Layers) FindMaxError(expect int) float64 {
	maxError := -math.MaxFloat64
	outputLayer := l.Neurons[l.LayersNum-1]

	for i := 0; i < outputLayer.Len(); i++ {
		currentError := 0.0
		if i == expect {
			currentError = 1 - outputLayer.AtVec(i)
		} else {
			currentError = outputLayer.AtVec(i)
		}
		if currentError > maxError {
			maxError = currentError
		}
	}

	return maxError
}

func NewLayers(sizes ...int) *Layers {
	layers := &Layers{
		LayersNum: LayersNum,
		Sizes:     sizes,
	}

	layers.InitWeights()
	layers.InitBias()
	layers.InitLayers()
	layers.InitLastChange()

	return layers
}
