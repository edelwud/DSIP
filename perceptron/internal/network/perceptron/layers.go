package perceptron

import "perceptron/internal/layers"

type Layers struct {
	Distribution layers.Layer
	Hidden       layers.Layer
	Output       layers.Layer
}
