package utils

import "testing"

const GrayscaleImagePath = "./resources/tests/image_grayscale.jpg"

func TestGrayscaleImage(t *testing.T) {
	img, err := ReadImageJpeg("./resources/tests/image2.jpg")
	if err != nil {
		t.Fatal(err)
	}

	grayImage := GrayscaleImage(img)

	err = WriteImageJpeg(grayImage, GrayscaleImagePath)
	if err != nil {
		t.Fatal(err)
	}
}
