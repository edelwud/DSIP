package network

type Network interface {
	BackPropagation(expect float64)
	ForwardFeed() float64
	UpdateWeights(lr float64)
}
