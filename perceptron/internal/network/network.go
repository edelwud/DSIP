package network

import "gonum.org/v1/gonum/mat"

type Network interface {
	Training(shapes []*mat.VecDense)
}
