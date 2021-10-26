package hopfield

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"hopfield/components/binarization"
	"hopfield/components/utils"
	"io/ioutil"
	"testing"
)

func TestGetWeights(t *testing.T) {
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

	W := GetWeights(images...)
	fmt.Println(mat.Formatted(W))
}
