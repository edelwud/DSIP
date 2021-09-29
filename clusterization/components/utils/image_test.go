package utils

import "testing"

const (
	ImagePath      = "./resources/tests/image.jpg"
	ImageStorePath = "./resources/tests/image_store.jpg"
)

func TestReadImageJpeg(t *testing.T) {
	_, err := ReadImageJpeg(ImagePath)
	if err != nil {
		t.Fatal(err)
	}
}

func TestWriteImageJpeg(t *testing.T) {
	img, err := ReadImageJpeg(ImagePath)
	if err != nil {
		t.Fatal(err)
	}

	err = WriteImageJpeg(img, ImageStorePath)
	if err != nil {
		t.Fatal(err)
	}
}
