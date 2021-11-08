package perceptron

import (
	"gonum.org/v1/gonum/mat"
	"math"
	"math/rand"
)

type Layers struct {
	LayersNum int
	Sizes     []int

	Neurons []*mat.VecDense
	Error   []*mat.VecDense
	Weights []*mat.Dense
	Bios    []*mat.VecDense
	BiosVal *mat.VecDense
}

func (l *Layers) InitLayers() {
	neurons := make([]*mat.VecDense, l.LayersNum)

	for i := range neurons {
		n := make([]float64, l.Sizes[i])
		neurons[i] = mat.NewVecDense(l.Sizes[i], n)
	}

	l.Neurons = neurons
}

func (l *Layers) InitWeights() {
	weights := make([]*mat.Dense, l.LayersNum-1)

	for i := range weights {
		w := make([]float64, l.Sizes[i+1]*l.Sizes[i])
		for j := range w {
			sign := rand.Float64()
			if sign > 0.5 {
				w[j] = rand.Float64()
			} else {
				w[j] = -rand.Float64()
			}
		}
		weights[i] = mat.NewDense(l.Sizes[i+1], l.Sizes[i], w)
	}

	l.Weights = weights
}

func (l *Layers) InitBios() {
	bios := make([]*mat.VecDense, l.LayersNum-1)

	for i := range bios {
		b := make([]float64, l.Sizes[i+1])
		for j := range b {
			sign := rand.Float64()
			if sign > 0.5 {
				b[j] = rand.Float64()
			} else {
				b[j] = -rand.Float64()
			}
		}
		bios[i] = mat.NewVecDense(l.Sizes[i+1], b)
	}

	bV := make([]float64, l.LayersNum-1)
	for i := range bV {
		bV[i] = 1
	}
	biosVal := mat.NewVecDense(l.LayersNum-1, bV)

	l.Bios = bios
	l.BiosVal = biosVal
}

func (l *Layers) InitErrors() {
	errors := make([]*mat.VecDense, l.LayersNum)

	for i := range errors {
		e := make([]float64, l.Sizes[i])
		errors[i] = mat.NewVecDense(len(e), e)
	}

	l.Error = errors
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

func NewLayers(layersNum int, sizes ...int) *Layers {
	layers := &Layers{
		LayersNum: layersNum,
		Sizes:     sizes,
	}

	layers.InitWeights()
	layers.InitBios()
	layers.InitLayers()
	layers.InitErrors()

	return layers
}
