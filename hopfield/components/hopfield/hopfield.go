package hopfield

import (
	"gonum.org/v1/gonum/mat"
	"image"
	"math/rand"
)

const MaxSyncHopfieldIterations = 1000

func SyncHopfield(img *image.Gray, golden []*image.Gray) *image.Gray {
	goldenVecs := make([]*mat.VecDense, 0)
	for _, object := range golden {
		vec := ConvertToVec(object)
		goldenVecs = append(goldenVecs, vec)
	}

	vec := ConvertToVec(img)
	W := GetWeights(goldenVecs...)
	m := mat.NewVecDense(vec.Len(), make([]float64, vec.Len()))

	i := 0

	for CompareWithGolden(vec, goldenVecs) != true {
		m.MulVec(W, vec)
		m = Activation(m)
		vec.CopyVec(m)

		if i >= MaxSyncHopfieldIterations {
			break
		}
		i++
	}

	return DenormalizeObject(img.Bounds().Dx(), img.Bounds().Dy(), vec)
}

func AsyncHopfield(img *image.Gray, golden []*image.Gray) *image.Gray {
	goldenVecs := make([]*mat.VecDense, 0)
	for _, object := range golden {
		vec := ConvertToVec(object)
		goldenVecs = append(goldenVecs, vec)
	}

	vec := ConvertToVec(img)
	W := GetWeights(goldenVecs...)
	m := mat.NewVecDense(vec.Len(), make([]float64, vec.Len()))

	for CompareWithGolden(vec, goldenVecs) != true {
		m.MulVec(W, vec)
		m = Activation(m)

		if Compare(vec, m) {
			return img
		}

		i := rand.Int() % vec.Len()
		vec.SetVec(i, m.AtVec(i))
	}

	return DenormalizeObject(img.Bounds().Dx(), img.Bounds().Dy(), vec)
}
