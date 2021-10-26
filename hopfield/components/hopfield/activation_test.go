package hopfield

import (
	"gonum.org/v1/gonum/mat"
	"hopfield/components/binarization"
	"hopfield/components/utils"
	"io/ioutil"
	"testing"
)

func TestSignFunction(t *testing.T) {
	value := SignFunction(-100)
	if value >= 0 {
		t.Fatal("unexpected value: ", value)
	}
}

func TestActivation(t *testing.T) {
	dir, err := ioutil.ReadDir("../../resources/training")
	if err != nil {
		t.Fatal(err)
	}

	images := make([]*mat.VecDense, 0)
	for _, filename := range dir {
		img, err := utils.ReadImagePNG("../../resources/training/" + filename.Name())
		if err != nil {
			t.Fatal(err)
		}

		gray := utils.GrayscaleImage(img)
		binRunner := binarization.CreateThresholdBinarization(gray, 128)
		normalized := NormalizeObject(binRunner.Process())
		vec := mat.NewVecDense(len(normalized), normalized)

		images = append(images, vec)
	}
}
