package sigma_activation

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"testing"
)

func TestSigmaActivation_Run(t *testing.T) {
	sigmaActivation := NewSigmaActivation()
	vec := mat.NewVecDense(3, []float64{1, 0, 0.2})
	result := sigmaActivation.Apply(vec)
	fmt.Println(mat.Formatted(result))
}

func TestSigmaActivation_Derivative(t *testing.T) {
	sigmaActivation := NewSigmaActivation()
	vec := mat.NewVecDense(3, []float64{1, 0, 0.2})
	result := sigmaActivation.Derivative(vec)
	fmt.Println(mat.Formatted(result))
}
