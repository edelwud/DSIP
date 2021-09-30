package binarization

import (
	"clusterization/components/utils"
	"testing"
)

const (
	ImageThresholdPath      = "./../../resources/easy/P0001460.jpg"
	ImageThresholdStorePath = "./../../resources/easy_output/image_threshold.jpg"
)

func TestThreshold_Process(t *testing.T) {
	image, err := utils.ReadImageJpeg(ImageThresholdPath)
	if err != nil {
		t.Fatal(err)
	}

	grayImage := utils.GrayscaleImage(image)
	binarization := CreateThresholdBinarization(grayImage, 200)
	processedImage := binarization.Process()

	err = utils.WriteImageJpeg(processedImage, ImageThresholdStorePath)
	if err != nil {
		t.Fatal(err)
	}
}
