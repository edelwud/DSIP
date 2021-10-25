package hopfield

import "gonum.org/v1/gonum/mat"

func Combinations(images ...*mat.Dense) [][]*mat.Dense {
	combinations := make([][]*mat.Dense, 0)

	for i := 0; i < len(images)-1; i++ {
		for j := i + 1; j < len(images); j++ {
			combinations = append(combinations, []*mat.Dense{
				images[i],
				images[j],
			})
		}
	}

	return combinations
}

func GetWeights(images ...*mat.Dense) *mat.Dense {
	r, c := images[0].Dims()
	result := mat.NewDense(r, c, make([]float64, r*c))

	for _, combination := range Combinations(images...) {
		mul := mat.NewDense(r, c, make([]float64, r*c))
		mul.Mul(combination[0].T(), combination[1])
		combination[0].T()
		result.Add(result, mul)
	}

	for i := 0; i < r; i++ {
		result.Set(i, i, 0)
	}

	return result
}
