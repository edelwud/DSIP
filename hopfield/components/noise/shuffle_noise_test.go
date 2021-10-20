package noise

import (
	"hopfield/components/binarization"
	"hopfield/components/utils"
	"testing"
)

const (
	ImagePath = "../../resources/training/train_1.png"
	ImageDest = "../../resources/temp/train_1_shuffled.png"
)

func TestShuffleNoise_Run(t *testing.T) {
	img, err := utils.ReadImagePNG(ImagePath)
	if err != nil {
		t.Fatal(err)
	}

	gray := utils.GrayscaleImage(img)
	threshold := binarization.CreateThresholdBinarization(gray, 128)
	bin := threshold.Process()

	shuffle := CreateShuffleNoise(bin, 0.10)
	shuffled := shuffle.Run()

	err = utils.WriteImagePNG(shuffled, ImageDest)
	if err != nil {
		t.Fatal(err)
	}
}
