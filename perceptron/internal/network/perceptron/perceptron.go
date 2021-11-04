package perceptron

import (
	"gonum.org/v1/gonum/mat"
	"perceptron/internal/activation"
	"perceptron/internal/layers"
	distributionLayer "perceptron/internal/layers/distribution_layer"
	hiddenLayer "perceptron/internal/layers/hidden_layer"
	outputLayer "perceptron/internal/layers/output_layer"
	"perceptron/internal/network"
)

type Perceptron struct {
	Activation   activation.Activation
	Shapes       []*mat.VecDense
	Layers       []layers.Layer
	NeuronsValue []float64
	NeuronsError []float64
	BiosValue    []float64
	BiosError    []float64
	Config       network.Config
}

func (p Perceptron) BackPropagation(expect float64) {
	return
}

func (p Perceptron) ForwardFeed() float64 {
	return 0
}

func (p Perceptron) UpdateWeights(lr float64) {
	return
}

func NewPerceptron(activation activation.Activation, config network.Config) network.Network {
	distribution := distributionLayer.NewDistributionLayer(config.DistributionLength)
	distribution.GenerateWeights(config.DistributionLength, config.HiddenLength)

	hidden := hiddenLayer.NewHiddenLayer(config.HiddenLength)
	hidden.GenerateWeights(config.HiddenLength, config.OutputLength)
	hidden.GenerateThreshold()

	output := outputLayer.NewOutputLayer(config.OutputLength)
	output.GenerateThreshold()

	return &Perceptron{
		Activation: activation,
		Layers: []layers.Layer{
			distribution,
			hidden,
			output,
		},
	}
}
