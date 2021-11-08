package perceptron

import (
	sigma "perceptron/internal/activation/sigma_activation"
	"perceptron/internal/network"
	"perceptron/internal/utils"
	"testing"
)

func TestPerceptron_Train(t *testing.T) {
	config := network.NewPerceptronDefaultConfig()
	activation := sigma.NewSigmaActivation()
	perceptron := NewPerceptron(activation, config)

	shapes, err := utils.ReadAllShapes("../../../resources/training")
	if err != nil {
		t.Fatal(err)
	}

	perceptron.Training(shapes)
}
