package utils

import "testing"

const (
	ImagePath      = "./resources/tests/image.jpg"
	ImageStorePath = "./resources/tests/image_store.jpg"
)

func TestJpegImageReading(t *testing.T) {
	image, err := ReadImageJpeg(ImagePath)
	if err != nil {
		t.Fatal(err)
	}

	if image != nil {
		t.Log("Image opened successfully")
	}
}

func TestJpegImageWriting(t *testing.T) {
	img, err := ReadImageJpeg(ImagePath)
	if err != nil {
		t.Fatal(err)
	}

	err = WriteImageJpeg(img, ImageStorePath)
	if err != nil {
		t.Fatal(err)
	}
}
