package cluster

import (
	"clusterization/components/binarization"
	"clusterization/components/blur"
	"clusterization/components/utils"
	"testing"
)

const (
	ImageArrangeShapes            = "./../../resources/easy/P0001461.jpg"
	ImageArrangeShapesGaussian    = "./../../resources/easy_output/image_arranged_shapes_gaussian.jpg"
	ImageArrangeShapesGaussianBin = "./../../resources/easy_output/image_arranged_shapes_gaussian_bin.jpg"
	ImageArrangeShapesStore       = "./../../resources/easy_output/image_arranged_shapes.jpg"
	ImageArrangeShapesTestStore   = "./../../resources/easy_output/image_arranged_shapes_test.jpg"
)

func TestArrangeShapes(t *testing.T) {
	image, err := utils.ReadImageJpeg(ImageArrangeShapes)
	if err != nil {
		t.Fatal(err)
	}

	gaussianBlur := blur.CreateGaussianBlur(image, 5)
	blurredImage := gaussianBlur.Process()
	err = utils.WriteImageJpeg(blurredImage, ImageArrangeShapesGaussian)
	if err != nil {
		t.Fatal(err)
	}

	grayImage := utils.GrayscaleImage(blurredImage)
	bin := binarization.CreateThresholdBinarization(grayImage, 200)
	binarizedImage := bin.Process()
	err = utils.WriteImageJpeg(binarizedImage, ImageArrangeShapesGaussianBin)
	if err != nil {
		t.Fatal(err)
	}

	contourMask := CreateContourAreaMask()
	connectedAreas := CreateConnectedAreasAnalyzer(bin.Process(), contourMask, 1)
	figures := connectedAreas.FindConnectedAreas()
	x, y := connectedAreas.Image.Bounds().Max.X, connectedAreas.Image.Bounds().Max.Y

	arrangedShapes := ArrangeShapes(figures, x, y)
	err = utils.WriteImageJpeg(arrangedShapes, ImageArrangeShapesStore)
	if err != nil {
		t.Fatal(err)
	}

	err = utils.WriteImageJpeg(connectedAreas.Image, ImageArrangeShapesTestStore)
	if err != nil {
		t.Fatal(err)
	}
}
