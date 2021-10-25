package hopfield

import (
	"gopkg.in/ffmt.v1"
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

func TestCombinations(t *testing.T) {
	dir, err := ioutil.ReadDir("../../resources/training")
	if err != nil {
		t.Fatal(err)
	}

	images := make([]Matrix, 0)
	for _, filename := range dir {
		img, err := utils.ReadImagePNG("../../resources/training/" + filename.Name())
		if err != nil {
			t.Fatal(err)
		}

		gray := utils.GrayscaleImage(img)
		binRunner := binarization.CreateThresholdBinarization(gray, 128)
		normalized := NormalizeObject(binRunner.Process())
		matrix := ConvertToMatrix(img.Bounds().Dx(), img.Bounds().Dy(), normalized)

		images = append(images, matrix)
	}

	combinations := Combinations(images...)
	ffmt.P(combinations)
}
