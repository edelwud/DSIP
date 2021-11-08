package perceptron

import (
	"gonum.org/v1/gonum/mat"
	"perceptron/internal/activation"
	"perceptron/internal/layers"
)

type Layers struct {
	Activation   activation.Activation
	Distribution layers.Layer
	Hidden       layers.Layer
	Output       layers.Layer
}

func (l Layers) CalculateHiddenLayerWeights() {
	l.CalculateLayerWeights(l.Hidden.N(), l.Hidden.Threshold(), l.Distribution.N(), l.Distribution.W())
}

func (l Layers) CalculateOutputLayerWeights() {
	l.CalculateLayerWeights(l.Output.N(), l.Output.Threshold(), l.Hidden.N(), l.Hidden.W())
}

func (l Layers) CalculateLayerWeights(first *mat.VecDense, threshold *mat.VecDense, second *mat.VecDense, weights *mat.Dense) {
	for i := 0; i < first.Len(); i++ {
		sum := 0.0
		for j := 0; j < second.Len(); j++ {
			sum += second.AtVec(j) * weights.At(j, i)
		}
		activated := l.Activation.ApplyValue(sum + threshold.AtVec(i))
		first.SetVec(i, activated)
	}
}

func (l Layers) RecalculateOutputLayerWeights(alpha float64, valid int) {
	hiddenOutputWeights := l.Hidden.W()
	r, c := hiddenOutputWeights.Dims()
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			outputState := l.Output.N().AtVec(j)
			hiddenState := l.Hidden.N().AtVec(i)
			currentWeight := hiddenOutputWeights.At(i, j)

			exposed := 0.0
			if j == valid {
				exposed = 1
			}

			newWeight := currentWeight + alpha*outputState*(1-outputState)*(outputState-exposed)*hiddenState
			hiddenOutputWeights.Set(i, j, newWeight)
		}
	}

	for i := 0; i < l.Output.N().Len(); i++ {
		outputState := l.Output.N().AtVec(i)
		outputThreshold := l.Output.Threshold()
		currentThreshold := outputThreshold.AtVec(i)

		exposed := 0.0
		if i == valid {
			exposed = 1
		}

		outputThreshold.SetVec(i, currentThreshold+alpha*outputState*(1-outputState)*(exposed-outputState))
	}
}

func (l Layers) RecalculateHiddenLayerWeights(alpha float64, valid int) {
	distributionHiddenWeights := l.Distribution.W()
	hiddenOutputWeights := l.Hidden.W()
	r, c := distributionHiddenWeights.Dims()
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			mistake := 0.0
			outputLayer := l.Output.N()
			for k := 0; k < outputLayer.Len(); k++ {
				outputState := outputLayer.AtVec(k)

				exposed := 0.0
				if i == valid {
					exposed = 1
				}

				mistake += (exposed - outputState) * outputState * (1 - outputState) * hiddenOutputWeights.At(j, k)
			}

			distributionLayer := l.Distribution.N()
			hideLayer := l.Hidden.N()
			temp := alpha * hideLayer.AtVec(j) * (1 - hideLayer.AtVec(j)) * mistake
			currentWeight := distributionHiddenWeights.At(i, j)
			distributionHiddenWeights.Set(i, j, currentWeight+temp*distributionLayer.AtVec(i))

			if i == 0 {
				hiddenThreshold := l.Hidden.Threshold()
				currentThreshold := hiddenThreshold.AtVec(j)
				hiddenThreshold.SetVec(j, currentThreshold+temp)
			}
		}
	}
}
