package hopfield

import (
	mat "github.com/gonum/matrix/mat64"
	"gopkg.in/ffmt.v1"
	"hopfield/components/binarization"
	"hopfield/components/utils"
	"io/ioutil"
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

func TestDenormalizeObject(t *testing.T) {
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

	noised, err := utils.ReadImagePNG("../../resources/shuffle/train_1_10.png")
	if err != nil {
		t.Fatal(err)
	}

	gray := utils.GrayscaleImage(noised)
	binRunner := binarization.CreateThresholdBinarization(gray, 128)
	normalized := NormalizeObject(binRunner.Process())
	noiseMatrix := mat.NewDense(noised.Bounds().Dy(), noised.Bounds().Dx(), normalized)

	m := AsyncHopfield(noiseMatrix, images)
	img := DenormalizeObject(m)
	err = utils.WriteImagePNG(img, "../../resources/results/train_1_10.png")
	if err != nil {
		t.Fatal(err)
	}
}

func TestConvertToMatrix(t *testing.T) {
	img, err := utils.ReadImagePNG("../../resources/training/train_1.png")
	if err != nil {
		t.Fatal(err)
	}

	gray := utils.GrayscaleImage(img)
	binRunner := binarization.CreateThresholdBinarization(gray, 128)

	data := NormalizeObject(binRunner.Process())
	matrix := mat.NewDense(img.Bounds().Dy(), img.Bounds().Dx(), data)

	_, err = ffmt.P(matrix)
	if err != nil {
		t.Fatal(err)
	}
}
