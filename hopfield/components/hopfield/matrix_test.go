package hopfield

import (
	ffmt "gopkg.in/ffmt.v1"
	"hopfield/components/binarization"
	"hopfield/components/utils"
	"testing"
)

func TestZeroingMainDiagonal(t *testing.T) {
	img, err := utils.ReadImagePNG("../../resources/training/train_1.png")
	if err != nil {
		t.Fatal(err)
	}

	gray := utils.GrayscaleImage(img)
	binRunner := binarization.CreateThresholdBinarization(gray, 128)

	data := NormalizeObject(binRunner.Process())
	matrix := ConvertToMatrix(img.Bounds().Dx(), img.Bounds().Dy(), data)
	matrix = ZeroingMainDiagonal(matrix)
	ffmt.P(matrix)
}

func TestTranspose(t *testing.T) {
	img, err := utils.ReadImagePNG("../../resources/training/train_1.png")
	if err != nil {
		t.Fatal(err)
	}

	gray := utils.GrayscaleImage(img)
	binRunner := binarization.CreateThresholdBinarization(gray, 128)

	data := NormalizeObject(binRunner.Process())
	matrix := ConvertToMatrix(img.Bounds().Dx(), img.Bounds().Dy(), data)
	transposedMatrix := Transpose(matrix)
	ffmt.P(transposedMatrix)
}
