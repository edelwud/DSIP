package hopfield

import (
	"gonum.org/v1/gonum/mat"
	"gopkg.in/ffmt.v1"
	"hopfield/components/binarization"
	"hopfield/components/utils"
	"io/ioutil"
	"testing"
)

func TestCombinations(t *testing.T) {
	dir, err := ioutil.ReadDir("../../resources/training")
	if err != nil {
		t.Fatal(err)
	}

	images := make([]*mat.Dense, 0)
	for _, filename := range dir {
		img, err := utils.ReadImagePNG("../../resources/training/" + filename.Name())
		if err != nil {
			t.Fatal(err)
		}

		gray := utils.GrayscaleImage(img)
		binRunner := binarization.CreateThresholdBinarization(gray, 128)
		normalized := NormalizeObject(binRunner.Process())
		matrix := mat.NewDense(img.Bounds().Dy(), img.Bounds().Dx(), normalized)

		images = append(images, matrix)
	}

	combinations := Combinations(images...)
	ffmt.P(combinations)
}

func TestGetWeights(t *testing.T) {
	dir, err := ioutil.ReadDir("../../resources/training")
	if err != nil {
		t.Fatal(err)
	}

	images := make([]*mat.Dense, 0)
	for _, filename := range dir {
		img, err := utils.ReadImagePNG("../../resources/training/" + filename.Name())
		if err != nil {
			t.Fatal(err)
		}

		gray := utils.GrayscaleImage(img)
		binRunner := binarization.CreateThresholdBinarization(gray, 128)
		normalized := NormalizeObject(binRunner.Process())
		matrix := mat.NewDense(img.Bounds().Dy(), img.Bounds().Dx(), normalized)

		images = append(images, matrix)
	}

	w := GetWeights(images...)
	ffmt.P(w)
}
