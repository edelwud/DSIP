package noise

import (
	"gonum.org/v1/gonum/mat"
)

type Noise interface {
	Run() *mat.VecDense
}
