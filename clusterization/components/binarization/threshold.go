package binarization

import (
	"image"
	"image/color"
)

// Threshold binarization implementation
type Threshold struct {
	Image     *image.Gray
	Threshold int
}

// Process runs Threshold binarization algorithm
func (t Threshold) Process() *image.Gray {
	img := image.NewGray(t.Image.Bounds())

	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			if t.Image.GrayAt(x, y).Y > uint8(t.Threshold) {
				img.Set(x, y, color.Gray{Y: 255})
			} else {
				img.Set(x, y, color.Gray{Y: 0})
			}
		}
	}

	return img
}

// CreateThresholdBinarization creates Threshold binarization exemplar
func CreateThresholdBinarization(img *image.Gray, threshold int) Binarization {
	return &Threshold{
		Image:     img,
		Threshold: threshold,
	}
}
