package perceptron

import (
	"gonum.org/v1/gonum/mat"
	"perceptron/internal/activation"
	distributionLayer "perceptron/internal/layers/distribution_layer"
	hiddenLayer "perceptron/internal/layers/hidden_layer"
	outputLayer "perceptron/internal/layers/output_layer"
	"perceptron/internal/network"
)

type Perceptron struct {
	Activation activation.Activation
	Shapes     []*mat.VecDense
	Layers     Layers
	Config     network.Config
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

func (p Perceptron) Train(shapes ...*mat.VecDense) {
	shapesTrained := make([]bool, len(shapes))
trainLoop:
	for {
		for i, shape := range shapes {
			if shapesTrained[i] == true {
				continue
			}
			p.Layers.Distribution.Fill(shape)
		}
		for _, value := range shapesTrained {
			if value == false {
				continue trainLoop
			}
		}
	}
}

func NewPerceptron(activation activation.Activation, config *network.Config) network.Network {
	distribution := distributionLayer.NewDistributionLayer(config.DistributionLength)
	distribution.GenerateWeights(config.DistributionLength, config.HiddenLength)

	hidden := hiddenLayer.NewHiddenLayer(config.HiddenLength)
	hidden.GenerateWeights(config.HiddenLength, config.OutputLength)
	hidden.GenerateThreshold()

	output := outputLayer.NewOutputLayer(config.OutputLength)
	output.GenerateThreshold()

	return &Perceptron{
		Activation: activation,
		Layers: Layers{
			distribution,
			hidden,
			output,
		},
	}
}
