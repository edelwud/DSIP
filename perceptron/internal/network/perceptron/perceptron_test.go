package perceptron

import (
	"fmt"
	"io/ioutil"
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

	dir, err := ioutil.ReadDir("../../../resources/shuffle")
	if err != nil {
		t.Fatal(err)
	}

	for _, file := range dir {
		if file.Name() == ".gitkeep" {
			continue
		}
		shape, err := utils.ReadShape("../../../resources/shuffle/" + file.Name())
		if err != nil {
			t.Fatal(err)
		}
		index := perceptron.Recognize(shape)
		fmt.Printf("File: %s, Class: %d\n", file.Name(), index+1)
	}
}
