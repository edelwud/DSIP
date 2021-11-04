package perceptron

import (
	"gonum.org/v1/gonum/mat"
	"perceptron/internal/layers"
	"perceptron/internal/network"
)

type Perceptron struct {
	Shapes []*mat.VecDense
}

func (p Perceptron) Layers() []layers.Layer {
	return nil
}

func (p Perceptron) Train() *mat.Dense {
	return nil
}

func NewPerceptron(shapes ...*mat.VecDense) network.Network {
	return &Perceptron{
		Shapes: shapes,
	}
}
