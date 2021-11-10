package rbf

import (
	"gonum.org/v1/gonum/mat"
	"rbf/internal/activation"
	"rbf/internal/network"
	"rbf/pkg/clusterization/kmeans"
)

type RBF struct {
	Activation activation.Activation
	Layers     *Layers
	Config     *network.Config
	Centers    []*mat.VecDense
}

const (
	Width  = 6
	Height = 6
)

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
	result := make([]*mat.VecDense, len(shapes))
	for i, shape := range shapes {
		result[i] = p.PrepareShape(shape)
	}

	return result
}

func (p RBF) PrepareShape(shape *mat.VecDense) *mat.VecDense {
	rbfShape := mat.NewDense(Height, Width, shape.RawVector().Data)
	var d []kmeans.Point

	rbfShape.Apply(func(i, j int, v float64) float64 {
		if v == 1 {
			d = append(d, kmeans.Point{
				Entry: []float64{float64(i), float64(j)},
			})
		}
		return v
	}, rbfShape)

	centroids := kmeans.KMeans(d, uint64(p.Config.Centers), 0.1)

	buffer := make([]float64, 0)
	for _, c := range centroids {
		buffer = append(buffer, c.Center.Entry[0])
		buffer = append(buffer, c.Center.Entry[1])
	}
	return mat.NewVecDense(len(buffer), buffer)
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
	rbfShapes := p.Prepare(shapes)

	maxError := 0.1
	trained := false
	trainedShapes := make([]bool, len(rbfShapes))

	for !trained {
		trained = true
		for i, shape := range rbfShapes {
			p.LoadShape(shape)
			p.ForwardFeed()
			currentError := p.Layers.FindMaxError(i)
			if currentError > maxError {
				p.BackPropagation(i)
				trainedShapes[i] = false
			} else {
				trainedShapes[i] = true
			}
		}

		for i := range trainedShapes {
			if trainedShapes[i] == false {
				trained = false
			}
		}
	}
}

func (p RBF) Recognize(shape *mat.VecDense) int {
	result := p.PrepareShape(shape)
	p.LoadShape(result)
	return p.ForwardFeed()
}

func (p RBF) LoadShape(shape *mat.VecDense) {
	p.Layers.Neurons[0].CopyVec(shape)
}

func NewRBF(activation activation.Activation, config *network.Config) network.Network {
	return &RBF{
		Activation: activation,
		Config:     config,
		Layers: NewLayers(
			config.Centers*2,
			config.OutputLength,
		),
	}
}
