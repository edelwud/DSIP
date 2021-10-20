package utils

import "testing"

const (
	ImageOpenPath      = "./../../resources/training/train_1.png"
	ImageOpenStorePath = "./../../resources/temp/train_1_store.png"
)

func TestReadImageJpeg(t *testing.T) {
	_, err := ReadImagePNG(ImageOpenPath)
	if err != nil {
		t.Fatal(err)
	}
}

func TestWriteImageJpeg(t *testing.T) {
	img, err := ReadImagePNG(ImageOpenPath)
	if err != nil {
		t.Fatal(err)
	}

	err = WriteImagePNG(img, ImageOpenStorePath)
	if err != nil {
		t.Fatal(err)
	}
}
