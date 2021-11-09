package noise

import (
	"gonum.org/v1/gonum/mat"
	"math/rand"
)

type ShuffleNoise struct {
	Intensity float64
	Shape     *mat.VecDense
}

func (n ShuffleNoise) Run() *mat.VecDense {
	shuffled := make([]float64, 0)

	for i := 0; i < n.Shape.Len(); i++ {
		pattern := n.Intensity - rand.Float64()
		if pattern >= 0 {
			shuffled = append(shuffled, 1-n.Shape.AtVec(i))
		} else {
			shuffled = append(shuffled, n.Shape.AtVec(i))
		}
	}

	return mat.NewVecDense(len(shuffled), shuffled)
}

func CreateShuffleNoise(shape *mat.VecDense, intensity float64) Noise {
	return &ShuffleNoise{
		Intensity: intensity,
		Shape:     shape,
	}
}
