package kohonen

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"io/ioutil"
	sigma "kohonen/internal/activation/sigma_activation"
	"kohonen/internal/network"
	"kohonen/internal/utils"
	"testing"
)

func TestKohonen_Train(t *testing.T) {
	config := network.NewKohonenDefaultConfig()
	activation := sigma.NewSigmaActivation()
	kohonen := NewKohonen(activation, config)

	shapes, err := utils.ReadAllShapes("../../../resources/training")
	if err != nil {
		t.Fatal(err)
	}

	kohonen.Training(shapes)

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
		index := kohonen.Recognize(shape)
		fmt.Printf("File: %s, Class: %d, ", file.Name(), index+1)
		fmt.Println(mat.Formatted(kohonen.OutputNeurons().T()))
	}
}
