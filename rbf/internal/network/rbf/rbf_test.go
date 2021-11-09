package rbf

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"io/ioutil"
	sigma "rbf/internal/activation/sigma_activation"
	"rbf/internal/network"
	"rbf/internal/utils"
	"testing"
)

func TestRBF_Train(t *testing.T) {
	config := network.NewRBFDefaultConfig()
	activation := sigma.NewSigmaActivation()
	rbf := NewRBF(activation, config)

	shapes, err := utils.ReadAllShapes("../../../resources/training")
	if err != nil {
		t.Fatal(err)
	}

	rbf.Training(shapes)

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
		index := rbf.Recognize(shape)
		fmt.Printf("File: %s, Class: %d, ", file.Name(), index+1)
		fmt.Println(mat.Formatted(rbf.OutputNeurons().T()))
	}
}
