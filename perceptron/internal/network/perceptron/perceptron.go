package perceptron

import (
	"gonum.org/v1/gonum/mat"
	"perceptron/internal/activation"
	"perceptron/internal/network"
)

type Perceptron struct {
	Activation activation.Activation
	Layers     *Layers
	Config     *network.Config
}

func (p Perceptron) BackPropagation(expect int) {
	outputLayer := p.Layers.Neurons[p.Layers.LayersNum-1]
	outputErrorLayer := p.Layers.Error[p.Layers.LayersNum-1]

	for i := 0; i < outputLayer.Len(); i++ {
		if i != expect {
			x := -outputLayer.AtVec(i) * p.Activation.DerivativeValue(outputLayer.AtVec(i))
			outputErrorLayer.SetVec(i, x)
		} else {
			x := (1.0 - outputLayer.AtVec(i)) * p.Activation.DerivativeValue(outputLayer.AtVec(i))
			outputErrorLayer.SetVec(i, x)
		}
	}

	neuronLayers := p.Layers.Neurons
	errorLayers := p.Layers.Error

	for i := p.Layers.LayersNum - 2; i > 0; i-- {
		errorLayers[i].MulVec(p.Layers.Weights[i].T(), errorLayers[i+1])
		for j := 0; j < p.Layers.Sizes[i]; j++ {
			x := errorLayers[i].AtVec(j) * p.Activation.DerivativeValue(neuronLayers[i].AtVec(j))
			errorLayers[i].SetVec(j, x)
		}
	}
}

func (p *Perceptron) ForwardFeed() int {
	for i := 1; i < p.Layers.LayersNum; i++ {
		p.Layers.Neurons[i].MulVec(p.Layers.Weights[i-1], p.Layers.Neurons[i-1])
		p.Layers.Neurons[i].AddVec(p.Layers.Neurons[i], p.Layers.Bios[i-1])
		p.Layers.Neurons[i].CopyVec(p.Activation.Apply(p.Layers.Neurons[i]))
	}

	return p.Layers.FindResult()
}

func (p Perceptron) UpdateWeights(lr float64) {
	for i := 0; i < p.Layers.LayersNum-1; i++ {
		for j := 0; j < p.Layers.Sizes[i+1]; j++ {
			for k := 0; k < p.Layers.Sizes[i]; k++ {
				x := p.Layers.Weights[i].At(j, k) + p.Layers.Neurons[i].AtVec(k)*p.Layers.Error[i+1].AtVec(j)*lr
				p.Layers.Weights[i].Set(j, k, x)
			}
		}
	}

	for i := 0; i < p.Layers.LayersNum-1; i++ {
		for k := 0; k < p.Layers.Sizes[i+1]; k++ {
			x := p.Layers.Bios[i].AtVec(k) + p.Layers.Error[i+1].AtVec(k)*lr
			p.Layers.Bios[i].SetVec(k, x)
		}
	}
}

func (p Perceptron) Training(shapes []*mat.VecDense) {
	ra := 0
	trained := false

	for !trained {
		trained = true
		ra = 0

		for i, shape := range shapes {
			p.LoadShape(shape)
			predict := p.ForwardFeed()

			if predict != i {
				p.BackPropagation(i)
				p.UpdateWeights(p.Config.Alpha)
			} else {
				ra++
			}

			if ra != len(shapes) {
				trained = false
			} else {
				trained = true
			}
		}
	}
}

func (p Perceptron) Recognize(shape *mat.VecDense) int {
	p.LoadShape(shape)
	return p.ForwardFeed()
}

func (p Perceptron) LoadShape(shape *mat.VecDense) {
	p.Layers.Neurons[0].CopyVec(shape)
}

func NewPerceptron(activation activation.Activation, config *network.Config) network.Network {
	return &Perceptron{
		Activation: activation,
		Config:     config,
		Layers: NewLayers(
			3,
			config.DistributionLength,
			config.HiddenLength,
			config.OutputLength,
		),
	}
}
