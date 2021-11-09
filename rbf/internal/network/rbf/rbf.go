package rbf

import (
	"gonum.org/v1/gonum/mat"
	"rbf/internal/activation"
	"rbf/internal/network"
)

type RBF struct {
	Activation activation.Activation
	Layers     *Layers
	Config     *network.Config
	Centers    int
}

func (p RBF) BackPropagation(expect int) {
	e := make([]float64, p.Config.OutputLength)
	e[expect] = 1
	errorVec := mat.NewVecDense(len(e), e)
	errorVec.SubVec(p.Layers.Neurons[1], errorVec)

	deltaVec := mat.NewVecDense(
		p.Layers.Sizes[1],
		make([]float64, p.Layers.Sizes[1]),
	)
	deltaVec.MulElemVec(p.Activation.Derivative(p.Layers.Neurons[1]), errorVec)

	deltaChange := mat.NewDense(
		p.Layers.Sizes[1],
		p.Layers.Sizes[0],
		make([]float64, p.Layers.Sizes[1]*p.Layers.Sizes[0]),
	)

	deltaChange.Mul(deltaVec, p.Layers.Neurons[0].T())
	deltaChange.Apply(func(_, _ int, v float64) float64 {
		return v * p.Config.DeltaRate
	}, deltaChange)

	lastDeltaChange := mat.NewDense(
		p.Layers.Sizes[1],
		p.Layers.Sizes[0],
		make([]float64, p.Layers.Sizes[1]*p.Layers.Sizes[0]),
	)
	lastDeltaChange.Copy(p.Layers.LastChange)
	lastDeltaChange.Apply(func(_, _ int, v float64) float64 {
		return v * p.Config.LastChangeRate
	}, lastDeltaChange)

	deltaChange.Add(deltaChange, lastDeltaChange)
	p.Layers.Weights.Sub(p.Layers.Weights, deltaChange)
	p.Layers.LastChange.Copy(lastDeltaChange)
}

func (p *RBF) ForwardFeed() int {
	p.Layers.Neurons[1].MulVec(p.Layers.Weights, p.Layers.Neurons[0])
	p.Layers.Neurons[1].AddVec(p.Layers.Neurons[1], p.Layers.Bias)
	p.Layers.Neurons[1].CopyVec(p.Activation.Apply(p.Layers.Neurons[1]))
	return p.Layers.FindResult()
}

func (p RBF) OutputNeurons() *mat.VecDense {
	return p.Layers.Neurons[p.Layers.LayersNum-1]
}

func (p RBF) Training(shapes []*mat.VecDense) {

	trainedShapes := make([]bool, len(shapes))
	trained := false
	maxError := 0.01

	for !trained {
		trained = true

		for i, shape := range shapes {
			p.LoadShape(shape)
			p.ForwardFeed()
			currentError := p.Layers.FindMaxError(i)

			if currentError > maxError {
				p.BackPropagation(i)
				trainedShapes[i] = false
			} else {
				trainedShapes[i] = true
			}
		}

		for i := range trainedShapes {
			if trainedShapes[i] == false {
				trained = false
			}
		}
	}
}

func (p RBF) Recognize(shape *mat.VecDense) int {
	p.LoadShape(shape)
	return p.ForwardFeed()
}

func (p RBF) LoadShape(shape *mat.VecDense) {
	p.Layers.Neurons[0].CopyVec(shape)
}

func NewRBF(activation activation.Activation, config *network.Config) network.Network {
	return &RBF{
		Activation: activation,
		Config:     config,
		Centers:    config.Centers,
		Layers: NewLayers(
			config.DistributionLength,
			config.OutputLength,
		),
	}
}
