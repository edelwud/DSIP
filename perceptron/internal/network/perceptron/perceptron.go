package perceptron

import (
	"gonum.org/v1/gonum/mat"
	"perceptron/internal/activation"
	"perceptron/internal/layers"
	"perceptron/internal/network"
)

type Perceptron struct {
	Activation activation.Activation
	Shapes     []*mat.VecDense
}

func (p Perceptron) Layers() []layers.Layer {
	return nil
}

func (p Perceptron) BackPropagation(expect float64) {
	return
}

func (p Perceptron) ForwardFeed() float64 {
	return 0
}

func (p Perceptron) WeightsUpdater(lr float64) {
	return
}

func NewPerceptron(activation activation.Activation, shapes ...*mat.VecDense) network.Network {
	return &Perceptron{
		Activation: activation,
		Shapes:     shapes,
	}
}
