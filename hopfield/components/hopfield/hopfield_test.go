package hopfield

import (
	"hopfield/components/binarization"
	"hopfield/components/utils"
	"image"
	"io/ioutil"
	"testing"
)

func TestAsyncHopfield(t *testing.T) {
	dir, err := ioutil.ReadDir("../../resources/training")
	if err != nil {
		t.Fatal(err)
	}

	images := make([]*image.Gray, 0)
	for _, file := range dir {
		img, err := utils.ReadImagePNG("../../resources/training/" + file.Name())
		if err != nil {
			t.Fatal(err)
		}

		gray := utils.GrayscaleImage(img)
		binRunner := binarization.CreateThresholdBinarization(gray, 128)

		images = append(images, binRunner.Process())
	}

	noised, err := utils.ReadImagePNG("../../resources/shuffle/train_1_10.png")
	if err != nil {
		t.Fatal(err)
	}

	gray := utils.GrayscaleImage(noised)
	binRunner := binarization.CreateThresholdBinarization(gray, 128)

	AsyncHopfield(binRunner.Process(), images)
}

func TestSyncHopfield(t *testing.T) {
	dir, err := ioutil.ReadDir("../../resources/training")
	if err != nil {
		t.Fatal(err)
	}

	images := make([]*image.Gray, 0)
	for _, file := range dir {
		img, err := utils.ReadImagePNG("../../resources/training/" + file.Name())
		if err != nil {
			t.Fatal(err)
		}

		gray := utils.GrayscaleImage(img)
		binRunner := binarization.CreateThresholdBinarization(gray, 128)

		images = append(images, binRunner.Process())
	}

	noised, err := utils.ReadImagePNG("../../resources/shuffle/train_1_10.png")
	if err != nil {
		t.Fatal(err)
	}

	gray := utils.GrayscaleImage(noised)
	binRunner := binarization.CreateThresholdBinarization(gray, 128)

	SyncHopfield(binRunner.Process(), images)
}
