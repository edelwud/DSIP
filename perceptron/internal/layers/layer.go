package layers

import "gonum.org/v1/gonum/mat"

type Layer interface {
	GenerateWeights(rows int, columns int)
	GenerateThreshold()
	Fill(vec *mat.VecDense)
}
