package blur

import (
	"clusterization/components/utils"
	"testing"
)

const (
	ImageGaussianBlurPath      = "./../../resources/easy/P0001460.jpg"
	ImageGaussianBlurStorePath = "./../../resources/easy_output/image_gaussian_blur.jpg"
)

func TestGaussianBlur_Process(t *testing.T) {
	image, err := utils.ReadImageJpeg(ImageGaussianBlurPath)
	if err != nil {
		t.Fatal(err)
	}

	gaussianBlur := CreateGaussianBlur(image, 3)
	blurredImg := gaussianBlur.Process()

	err = utils.WriteImageJpeg(blurredImg, ImageGaussianBlurStorePath)
	if err != nil {
		t.Fatal(err)
	}
}
