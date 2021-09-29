package binarization

import (
	"clusterization/components/utils"
	"testing"
)

const OtsuImageStorePath = "./resources/tests/image_otsu.jpg"

func TestOtsu_Process(t *testing.T) {
	image, err := utils.ReadImageJpeg(ImagePath)
	if err != nil {
		t.Fatal(err)
	}

	grayImage := utils.GrayscaleImage(image)
	binarization := CreateOtsuBinarization(grayImage)
	processedImage := binarization.Process()
	err = utils.WriteImageJpeg(processedImage, OtsuImageStorePath)
	if err != nil {
		t.Fatal(err)
	}
}
