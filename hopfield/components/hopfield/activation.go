package hopfield

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"math/rand"
)

func SignFunction(value float64) (result float64) {
	if value > 0 {
		result = 1
	} else {
		result = -1
	}
	return
}

func Activation(m *mat.Dense) *mat.Dense {
	r, c := m.Dims()
	result := mat.NewDense(r, c, make([]float64, r*c))

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			result.Set(i, j, SignFunction(m.At(i, j)))
		}
	}

	return result
}

func CompareWithGolden(img *mat.Dense, golden []*mat.Dense) bool {
	for _, gold := range golden {
		if Compare(gold, img) {
			fmt.Println(gold, img)
			return true
		}
	}
	return false
}

func Compare(m1 *mat.Dense, m2 *mat.Dense) bool {
	r, c := m1.Dims()

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if m1.At(i, j) != m2.At(i, j) {
				return false
			}
		}
	}

	return true
}

func SyncHopfield(img *mat.Dense, golden []*mat.Dense) {
	W := GetWeights(golden...)

	for CompareWithGolden(img, golden) != true {
		img.Mul(img, W)
		img = Activation(img)
	}
}

func AsyncHopfield(img *mat.Dense, golden []*mat.Dense) {
	W := GetWeights(golden...)

	r, c := img.Dims()
	m := mat.NewDense(r, c, make([]float64, r*c))

	for {
		m.Mul(img, W)
		m = Activation(m)

		if Compare(m, img) {
			break
		}

		for k := 0; k < 10; k++ {
			i := rand.Int() % r
			j := rand.Int() % c

			img.Set(i, j, m.At(i, j))
		}
	}

	if CompareWithGolden(img, golden) {
		println("kekw")
	}
}
