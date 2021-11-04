package network

import (
	"gonum.org/v1/gonum/mat"
	"perceptron/internal/layers"
)

type Network interface {
	Layers() []layers.Layer
	Train() *mat.Dense
}
