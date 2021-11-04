package layers

type Layer interface {
	GenerateWeights(rows int, columns int)
	GenerateThreshold()
}
