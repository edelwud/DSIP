package noise

import (
	"hopfield/components/binarization"
	"hopfield/components/utils"
	"strconv"
	"testing"
)

const (
	ImagePath = "../../resources/training/train_2.png"
)

func TestShuffleNoise_Run(t *testing.T) {
	img, err := utils.ReadImagePNG(ImagePath)
	if err != nil {
		t.Fatal(err)
	}

	gray := utils.GrayscaleImage(img)
	threshold := binarization.CreateThresholdBinarization(gray, 128)
	bin := threshold.Process()

	for i := 0.0; i <= 1; i += 0.1 {
		shuffle := CreateShuffleNoise(bin, i)
		shuffled := shuffle.Run()

		percent := int(i * 100)
		err = utils.WriteImagePNG(shuffled, "../../resources/shuffle/train_2_"+strconv.Itoa(percent)+".png")
		if err != nil {
			t.Fatal(err)
		}
	}
}
