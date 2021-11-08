package perceptron

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"math"
	"os"
	"perceptron/internal/activation"
	"perceptron/internal/layers"
	distributionLayer "perceptron/internal/layers/distribution_layer"
	hiddenLayer "perceptron/internal/layers/hidden_layer"
	outputLayer "perceptron/internal/layers/output_layer"
	"perceptron/internal/network"
)

type Perceptron struct {
	Activation activation.Activation
	Layers     Layers
	Config     *network.Config
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

func (p Perceptron) Train(shapes ...*mat.VecDense) error {
	shapesTrained := make([]bool, len(shapes))
	m := 0
trainLoop:
	for {
		for i := range shapes {
			if shapesTrained[i] == true {
				continue
			}
			p.Layers.Distribution.Fill(shapes[i])

			p.Layers.CalculateHiddenLayerWeights()
			p.Layers.CalculateOutputLayerWeights()

			mistake := p.FindMaxMistake(i, p.Layers.Output)
			if p.Config.Epsilon >= mistake {
				shapesTrained[i] = true
				continue
			}

			if m == 100000 {
				println(mistake)
				m = 0
			}

			p.Layers.RecalculateOutputLayerWeights(p.Config.Alpha, i)
			p.Layers.RecalculateHiddenLayerWeights(p.Config.Alpha, i)
			m++
		}
		for _, value := range shapesTrained {
			if value == false {
				continue trainLoop
			}
		}
		break
	}

	hiddenLayerFile, err := os.OpenFile("../../../resources/data/hidden.dat", os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		return err
	}

	hiddenFmt := mat.Formatted(p.Layers.Hidden.W())
	_, err = fmt.Fprint(hiddenLayerFile, hiddenFmt)
	if err != nil {
		return err
	}

	outputLayerFile, err := os.OpenFile("../../../resources/data/output.dat", os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		return err
	}

	outputFmt := mat.Formatted(p.Layers.Output.W())
	_, err = fmt.Fprint(outputLayerFile, outputFmt)
	if err != nil {
		return err
	}

	return nil
}

func (p Perceptron) FindMaxMistake(index int, output layers.Layer) float64 {
	maxMistake := 0.0
	for i := 0; i < output.N().Len(); i++ {
		delta := 0.0
		if i == index {
			delta = 1
		}
		delta = math.Abs(delta - output.N().AtVec(i))

		if maxMistake < delta {
			maxMistake = delta
		}
	}

	return maxMistake
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
		Config: config,
		Layers: Layers{
			activation,
			distribution,
			hidden,
			output,
		},
	}
}
