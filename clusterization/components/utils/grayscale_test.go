package utils

import "testing"

const (
	ImageGrayscalePath      = "./../../resources/easy/P0001460.jpg"
	ImageGrayscaleStorePath = "./../../resources/easy_output/image_grayscale.jpg"
)

func TestGrayscaleImage(t *testing.T) {
	img, err := ReadImageJpeg(ImageGrayscalePath)
	if err != nil {
		t.Fatal(err)
	}

	grayImage := GrayscaleImage(img)

	err = WriteImageJpeg(grayImage, ImageGrayscaleStorePath)
	if err != nil {
		t.Fatal(err)
	}
}
