package utils

import "testing"

const (
	ImageGrayscalePath      = "./../../resources/training/train_1.png"
	ImageGrayscaleStorePath = "./../../resources/temp/train_1_gray.png"
)

func TestGrayscaleImage(t *testing.T) {
	img, err := ReadImagePNG(ImageGrayscalePath)
	if err != nil {
		t.Fatal(err)
	}

	grayImage := GrayscaleImage(img)

	err = WriteImagePNG(grayImage, ImageGrayscaleStorePath)
	if err != nil {
		t.Fatal(err)
	}
}
