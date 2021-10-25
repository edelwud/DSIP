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

func TestMultiplication(t *testing.T) {
	img, err := utils.ReadImagePNG("../../resources/training/train_1.png")
	if err != nil {
		t.Fatal(err)
	}

	gray := utils.GrayscaleImage(img)
	binRunner := binarization.CreateThresholdBinarization(gray, 128)

	data := NormalizeObject(binRunner.Process())
	m1 := ConvertToMatrix(img.Bounds().Dx(), img.Bounds().Dy(), data)
	m2 := ConvertToMatrix(img.Bounds().Dx(), img.Bounds().Dy(), data)

	m3 := Multiplication(m1, m2)
	ffmt.P(m3)
}

func TestSum(t *testing.T) {
	img, err := utils.ReadImagePNG("../../resources/training/train_1.png")
	if err != nil {
		t.Fatal(err)
	}

	gray := utils.GrayscaleImage(img)
	binRunner := binarization.CreateThresholdBinarization(gray, 128)

	data := NormalizeObject(binRunner.Process())
	m1 := ConvertToMatrix(img.Bounds().Dx(), img.Bounds().Dy(), data)
	m2 := ConvertToMatrix(img.Bounds().Dx(), img.Bounds().Dy(), data)
	m3 := Sum(m1, m2)
	ffmt.P(m3)
}

func TestCompare(t *testing.T) {
	img, err := utils.ReadImagePNG("../../resources/training/train_1.png")
	if err != nil {
		t.Fatal(err)
	}

	gray := utils.GrayscaleImage(img)
	binRunner := binarization.CreateThresholdBinarization(gray, 128)

	data := NormalizeObject(binRunner.Process())
	m1 := ConvertToMatrix(img.Bounds().Dx(), img.Bounds().Dy(), data)
	m2 := ConvertToMatrix(img.Bounds().Dx(), img.Bounds().Dy(), data)
	if Compare(m1, m2) == false {
		t.Fatal("unexpected behavior")
	}
}
