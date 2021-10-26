package hopfield

import (
	"hopfield/components/binarization"
	"hopfield/components/utils"
	"testing"
)

func TestNormalizeObject(t *testing.T) {
	img, err := utils.ReadImagePNG("../../resources/training/train_1.png")
	if err != nil {
		t.Fatal(err)
	}

	gray := utils.GrayscaleImage(img)
	binRunner := binarization.CreateThresholdBinarization(gray, 128)

	data := NormalizeObject(binRunner.Process())

	for _, value := range data {
		if value == PositiveNormalization || value == NegativeNormalization {
			continue
		}
		t.Fatal("unexpected result")
	}
}
