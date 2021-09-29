package utils

import "testing"

const (
	ImageOpenPath      = "./../../resources/easy/P0001460.jpg"
	ImageOpenStorePath = "./../../resources/easy_output/image_store.jpg"
)

func TestReadImageJpeg(t *testing.T) {
	_, err := ReadImageJpeg(ImageOpenPath)
	if err != nil {
		t.Fatal(err)
	}
}

func TestWriteImageJpeg(t *testing.T) {
	img, err := ReadImageJpeg(ImageOpenPath)
	if err != nil {
		t.Fatal(err)
	}

	err = WriteImageJpeg(img, ImageOpenStorePath)
	if err != nil {
		t.Fatal(err)
	}
}
