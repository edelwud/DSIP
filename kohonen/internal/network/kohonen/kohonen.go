package kohonen

import (
	"gonum.org/v1/gonum/mat"
	"kohonen/internal/activation"
	"kohonen/internal/network"
	"math"
)

type Kohonen struct {
	Activation activation.Activation
	Layers     *Layers
	Config     *network.Config
	Centers    []*mat.VecDense
}

func (k Kohonen) BackPropagation(winnerIndex int) {
	r, _ := k.Layers.Weights.Dims()
	for i := 0; i < r; i++ {
		value := k.Layers.Weights.At(i, winnerIndex)
		result := value + k.Config.Beta*(k.Layers.Neurons[0].AtVec(i)-value)
		k.Layers.Weights.Set(i, winnerIndex, result)
	}
}

func (k *Kohonen) ForwardFeed() int {
	winnerNeuronIndex := 0
	winnerValue := -1.0

	for i := 0; i < k.Layers.Neurons[1].Len(); i++ {
		distance := 0.0
		for j := 0; j < k.Layers.Neurons[0].Len(); j++ {
			distance += math.Pow(k.Layers.Neurons[0].AtVec(j)-k.Layers.Weights.At(j, i), 2)
		}
		distance = math.Sqrt(distance)
		temp := math.Abs(distance) * k.Layers.WinningCount.AtVec(i)

		if winnerValue > temp || i == 0 {
			winnerValue = temp
			winnerNeuronIndex = i
		}
	}

	current := k.Layers.WinningCount.AtVec(winnerNeuronIndex)
	k.Layers.WinningCount.SetVec(winnerNeuronIndex, current+1)
	return winnerNeuronIndex
}

func (k Kohonen) OutputNeurons() *mat.VecDense {
	for i := 0; i < k.Layers.Neurons[1].Len(); i++ {
		inputWeightSum := 0.0
		for j := 0; j < k.Layers.Neurons[0].Len(); j++ {
			inputWeightSum += k.Layers.Neurons[0].AtVec(j) * k.Layers.Weights.At(j, i)
		}
		k.Layers.Neurons[1].SetVec(i, inputWeightSum)
	}

	return k.Layers.Neurons[k.Layers.LayersNum-1]
}

func (k *Kohonen) Training(shapes []*mat.VecDense) {
	maxError := 0.001
	trained := false
	trainedShapes := make([]bool, len(shapes))

	for !trained {
		trained = true
		for i, shape := range shapes {
			k.LoadShape(shape)
			winnerIndex := k.ForwardFeed()
			currentError := k.Layers.FindDistance(winnerIndex)
			if currentError > maxError {
				k.BackPropagation(i)
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

func (k Kohonen) Recognize(shape *mat.VecDense) int {
	k.LoadShape(shape)
	return k.ForwardFeed()
}

func (k Kohonen) LoadShape(shape *mat.VecDense) {
	k.Layers.Neurons[0].CopyVec(shape)
}

func NewKohonen(activation activation.Activation, config *network.Config) network.Network {
	return &Kohonen{
		Activation: activation,
		Config:     config,
		Layers: NewLayers(
			config.DistributionLength,
			config.OutputLength,
		),
	}
}
