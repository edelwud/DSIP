package utils

import (
	"image"
	"image/png"
	"os"
)

// ReadImagePNG reads png image from file
func ReadImagePNG(path string) (image.Image, error) {
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

	img, err := png.Decode(file)
	if err != nil {
		return nil, err
	}

	return img, nil
}

// WriteImagePNG writes png image to file
func WriteImagePNG(img image.Image, path string) error {
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

	err = png.Encode(file, img)
	if err != nil {
		return err
	}

	return nil
}
