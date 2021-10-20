package binarization

import (
	"hopfield/components/utils"
	"testing"
)

const (
	ImageThresholdPath      = "./../../resources/training/train_1.png"
	ImageThresholdStorePath = "./../../resources/temp/train_1_bin_threshold.jpg"
)

func TestThreshold_Process(t *testing.T) {
	image, err := utils.ReadImagePNG(ImageThresholdPath)
	if err != nil {
		t.Fatal(err)
	}

	grayImage := utils.GrayscaleImage(image)
	binarization := CreateThresholdBinarization(grayImage, 128)
	processedImage := binarization.Process()

	err = utils.WriteImagePNG(processedImage, ImageThresholdStorePath)
	if err != nil {
		t.Fatal(err)
	}
}
