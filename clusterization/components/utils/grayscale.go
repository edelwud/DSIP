package utils

import "image"

func GrayscaleImage(img image.Image) *image.Gray {
	grayImg := image.NewGray(img.Bounds())

	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			grayImg.Set(x, y, img.At(x, y))
		}
	}

	return grayImg
}
