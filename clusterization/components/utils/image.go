package utils

import (
	"image"
	"image/jpeg"
	"os"
)

// ReadImageJpeg reads jpeg image from file
func ReadImageJpeg(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)

	img, err := jpeg.Decode(file)
	if err != nil {
		return nil, err
	}

	return img, nil
}

// WriteImageJpeg writes jpeg image to file
func WriteImageJpeg(img image.Image, path string) error {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)

	err = jpeg.Encode(file, img, nil)
	if err != nil {
		return err
	}

	return nil
}
