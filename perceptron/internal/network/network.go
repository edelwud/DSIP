package network

import (
	"perceptron/internal/layers"
)

type Network interface {
	Layers() []layers.Layer
	BackPropagation(expect float64)
	ForwardFeed() float64
	WeightsUpdater(lr float64)
}
