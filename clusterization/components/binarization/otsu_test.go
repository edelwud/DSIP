package binarization

import (
	"clusterization/components/utils"
	"os"
	"testing"
)

const (
	ImageOtsuPath      = "./../../resources/easy/P0001460.jpg"
	ImageOtsuStorePath = "./../../resources/easy_output/image_otsu.jpg"
)

func TestOtsu_Process(t *testing.T) {
	t.Log(os.Getwd())

	image, err := utils.ReadImageJpeg(ImageOtsuPath)
	if err != nil {
		t.Fatal(err)
	}

	grayImage := utils.GrayscaleImage(image)
	binarization := CreateOtsuBinarization(grayImage)
	processedImage := binarization.Process()

	err = utils.WriteImageJpeg(processedImage, ImageOtsuStorePath)
	if err != nil {
		t.Fatal(err)
	}
}
