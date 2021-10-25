package hopfield

import (
	"gonum.org/v1/gonum/mat"
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

func TestActivation(t *testing.T) {
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

	W := GetWeights(images...)
	activated := Activation(W)
	ffmt.P(activated)
}

func TestAsyncHopfield(t *testing.T) {
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

	noised, err := utils.ReadImagePNG("../../resources/shuffle/train_2_30.png")
	if err != nil {
		t.Fatal(err)
	}

	gray := utils.GrayscaleImage(noised)
	binRunner := binarization.CreateThresholdBinarization(gray, 128)
	normalized := NormalizeObject(binRunner.Process())
	noiseMatrix := mat.NewDense(noised.Bounds().Dy(), noised.Bounds().Dx(), normalized)

	AsyncHopfield(noiseMatrix, images)
}

func TestSyncHopfield(t *testing.T) {
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

	noised, err := utils.ReadImagePNG("../../resources/shuffle/train_2_30.png")
	if err != nil {
		t.Fatal(err)
	}

	gray := utils.GrayscaleImage(noised)
	binRunner := binarization.CreateThresholdBinarization(gray, 128)
	normalized := NormalizeObject(binRunner.Process())
	noiseMatrix := mat.NewDense(noised.Bounds().Dy(), noised.Bounds().Dx(), normalized)

	SyncHopfield(noiseMatrix, images)
}
