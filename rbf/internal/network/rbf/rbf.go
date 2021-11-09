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
}

func (p RBF) BackPropagation(expect int) {
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

func (p *RBF) ForwardFeed() int {
	for i := 1; i < p.Layers.LayersNum; i++ {
		p.Layers.Neurons[i].MulVec(p.Layers.Weights[i-1], p.Layers.Neurons[i-1])
		p.Layers.Neurons[i].AddVec(p.Layers.Neurons[i], p.Layers.Bios[i-1])
		p.Layers.Neurons[i].CopyVec(p.Activation.Apply(p.Layers.Neurons[i]))
	}

	return p.Layers.FindResult()
}

func (p RBF) UpdateWeights(lr float64) {
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
				p.UpdateWeights(p.Config.Alpha)
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
		Layers: NewLayers(
			3,
			config.DistributionLength,
			config.HiddenLength,
			config.OutputLength,
		),
	}
}
