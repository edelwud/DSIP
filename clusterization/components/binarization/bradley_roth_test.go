package binarization

import (
	"clusterization/components/utils"
	"testing"
)

const (
	ImagePath      = "./resources/tests/image.jpg"
	ImageStorePath = "./resources/tests/image_bradley_roth.jpg"
)

func TestBradleyRoth_Process(t *testing.T) {
	image, err := utils.ReadImageJpeg(ImagePath)
	if err != nil {
		t.Fatal(err)
	}

	grayImage := utils.GrayscaleImage(image)
	binarization := CreateBradleyRothBinarization(grayImage, 0.15)
	processedImage := binarization.Process()
	err = utils.WriteImageJpeg(processedImage, ImageStorePath)
	if err != nil {
		t.Fatal(err)
	}
}
