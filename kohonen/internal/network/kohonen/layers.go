package kohonen

import (
	"gonum.org/v1/gonum/mat"
	"math"
	"math/rand"
)

type Layers struct {
	LayersNum int
	Sizes     []int

	Neurons      []*mat.VecDense
	Weights      *mat.Dense
	WinningCount *mat.VecDense
	Bias         *mat.VecDense
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
	l.Weights = mat.NewDense(l.Sizes[0], l.Sizes[1], w)
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

func (l *Layers) InitWinningCount() {
	win := make([]float64, l.Sizes[1])
	l.WinningCount = mat.NewVecDense(l.Sizes[1], win)
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

func (l Layers) FindDistance(winnerIndex int) float64 {
	distance := 0.0

	for i := 0; i < l.Neurons[0].Len(); i++ {
		distance += math.Pow(l.Neurons[0].AtVec(i)-l.Weights.At(i, winnerIndex), 2)
	}

	return math.Sqrt(distance)
}

func NewLayers(sizes ...int) *Layers {
	layers := &Layers{
		LayersNum: LayersNum,
		Sizes:     sizes,
	}

	layers.InitWeights()
	layers.InitBias()
	layers.InitLayers()
	layers.InitWinningCount()

	return layers
}
