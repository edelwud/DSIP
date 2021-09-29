package binarization

import (
	"clusterization/components/utils"
	"testing"
)

const (
	ImageBradleyRothPath      = "./../../resources/easy/P0001460.jpg"
	ImageBradleyRothStorePath = "./../../resources/easy_output/image_bradley_roth.jpg"
)

func TestBradleyRoth_Process(t *testing.T) {
	image, err := utils.ReadImageJpeg(ImageBradleyRothPath)
	if err != nil {
		t.Fatal(err)
	}

	grayImage := utils.GrayscaleImage(image)
	binarization := CreateBradleyRothBinarization(grayImage, 0.15)
	processedImage := binarization.Process()

	err = utils.WriteImageJpeg(processedImage, ImageBradleyRothStorePath)
	if err != nil {
		t.Fatal(err)
	}
}
