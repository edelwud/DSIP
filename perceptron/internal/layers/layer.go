package layers

import "gonum.org/v1/gonum/mat"

type Layer interface {
	W() *mat.Dense
	N() *mat.VecDense
	Threshold() *mat.VecDense

	GenerateWeights(rows int, columns int)
	GenerateThreshold()
	Fill(vec *mat.VecDense)
}
