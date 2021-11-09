package rbf

import (
	"gonum.org/v1/gonum/mat"
	"math"
	"math/rand"
	"rbf/internal/activation"
	"rbf/internal/network"
	"time"
)

type RBF struct {
	Activation activation.Activation
	Layers     *Layers
	Config     *network.Config
	Centers    []*mat.VecDense
}

func (p RBF) BackPropagation(expect int) {
	e := make([]float64, p.Config.OutputLength)
	e[expect] = 1
	errorVec := mat.NewVecDense(len(e), e)
	errorVec.SubVec(p.Layers.Neurons[1], errorVec)

	deltaVec := mat.NewVecDense(
		p.Layers.Sizes[1],
		make([]float64, p.Layers.Sizes[1]),
	)
	deltaVec.MulElemVec(p.Activation.Derivative(p.Layers.Neurons[1]), errorVec)

	deltaChange := mat.NewDense(
		p.Layers.Sizes[1],
		p.Layers.Sizes[0],
		make([]float64, p.Layers.Sizes[1]*p.Layers.Sizes[0]),
	)

	deltaChange.Mul(deltaVec, p.Layers.Neurons[0].T())
	deltaChange.Apply(func(_, _ int, v float64) float64 {
		return v * p.Config.DeltaRate
	}, deltaChange)

	lastDeltaChange := mat.NewDense(
		p.Layers.Sizes[1],
		p.Layers.Sizes[0],
		make([]float64, p.Layers.Sizes[1]*p.Layers.Sizes[0]),
	)
	lastDeltaChange.Copy(p.Layers.LastChange)
	lastDeltaChange.Apply(func(_, _ int, v float64) float64 {
		return v * p.Config.LastChangeRate
	}, lastDeltaChange)

	deltaChange.Add(deltaChange, lastDeltaChange)
	p.Layers.Weights.Sub(p.Layers.Weights, deltaChange)
	p.Layers.LastChange.Copy(lastDeltaChange)
}

func (p RBF) Prepare(shapes []*mat.VecDense) []*mat.VecDense {
	rbfShapes := make([]*mat.VecDense, len(shapes))
	for i, shape := range shapes {
		result := p.Gaussian(shape)
		rbfShapes[i] = mat.NewVecDense(result.Len(), make([]float64, result.Len()))
		rbfShapes[i].CopyVec(result)
	}

	return rbfShapes
}

func (p *RBF) CalculateCenters(shapes []*mat.VecDense) {
	p.Centers = make([]*mat.VecDense, p.Config.Centers)
	for i := range p.Centers {
		idx := genRandomIdx(len(shapes))
		p.Centers[i] = mat.NewVecDense(shapes[idx[i]].Len(), make([]float64, shapes[idx[i]].Len()))
		p.Centers[i].CopyVec(shapes[idx[i]])
	}
}

func (p RBF) Gaussian(shape *mat.VecDense) *mat.VecDense {
	result := mat.NewVecDense(p.Config.Centers, make([]float64, p.Config.Centers))
	div := 0.0
	for j := 0; j < p.Config.Centers; j++ {
		sum := 0.0
		for i := 0; i < p.Config.DistributionLength; i++ {
			delta := shape.AtVec(i) - p.Centers[j].AtVec(i)
			sum += delta * delta
		}
		result.SetVec(j, math.Exp(-8*sum))
		div += result.AtVec(j)
	}

	for i := 0; i < result.Len(); i++ {
		x := result.AtVec(i)
		result.SetVec(i, x/div)
	}

	return result
}

func genRandomIdx(N int) []int {
	A := make([]int, N)
	for i := 0; i < N; i++ {
		A[i] = i
	}

	for i := 0; i < N; i++ {
		j := i + int(rand.Float64()*float64(N-i))
		A[i], A[j] = A[j], A[i]
	}
	return A
}

func (p *RBF) ForwardFeed() int {
	p.Layers.Neurons[1].MulVec(p.Layers.Weights, p.Layers.Neurons[0])
	p.Layers.Neurons[1].AddVec(p.Layers.Neurons[1], p.Layers.Bias)
	p.Layers.Neurons[1].CopyVec(p.Activation.Apply(p.Layers.Neurons[1]))
	return p.Layers.FindResult()
}

func (p RBF) OutputNeurons() *mat.VecDense {
	return p.Layers.Neurons[p.Layers.LayersNum-1]
}

func (p *RBF) Training(shapes []*mat.VecDense) {
	p.CalculateCenters(shapes)
	rbfShapes := p.Prepare(shapes)
	maxIterations := 10000
	i := 0

	for i < maxIterations {
		for i, shape := range rbfShapes {
			p.LoadShape(shape)
			p.ForwardFeed()
			p.BackPropagation(i)
		}
		i++
	}
}

func (p RBF) Recognize(shape *mat.VecDense) int {
	result := p.Gaussian(shape)
	p.LoadShape(result)
	return p.ForwardFeed()
}

func (p RBF) LoadShape(shape *mat.VecDense) {
	p.Layers.Neurons[0].CopyVec(shape)
}

func NewRBF(activation activation.Activation, config *network.Config) network.Network {
	rand.Seed(time.Now().UnixNano())
	return &RBF{
		Activation: activation,
		Config:     config,
		Layers: NewLayers(
			config.Centers,
			config.OutputLength,
		),
	}
}
