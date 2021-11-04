package network

import "gonum.org/v1/gonum/mat"

type Network interface {
	BackPropagation(expect float64)
	ForwardFeed() float64
	UpdateWeights(lr float64)
	Train(shapes ...*mat.VecDense)
}
