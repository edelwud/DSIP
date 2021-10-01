package utils

import (
	"clusterization/components/binarization"
	"clusterization/components/cluster"
	"testing"
)

const (
	ImageArrangeShapes      = "./../../resources/easy/P0001460.jpg"
	ImageArrangeShapesStore = "./../../resources/easy_output/image_arranged_shapes.jpg"
)

func TestConnectedAreas_FindConnectedAreas(t *testing.T) {
	image, err := ReadImageJpeg(ImageArrangeShapes)
	if err != nil {
		t.Fatal(err)
	}

	grayImage := GrayscaleImage(image)
	bin := binarization.CreateThresholdBinarization(grayImage, 200)

	connectedAreas := cluster.CreateConnectedAreasAnalyzer(bin.Process(), 30)
	figures := connectedAreas.FindConnectedAreas()

	x, y := connectedAreas.Image.Bounds().Max.X, connectedAreas.Image.Bounds().Max.Y

	arrangedShapes := ArrangeShapes(figures, x, y)
	err = WriteImageJpeg(arrangedShapes, ImageArrangeShapesStore)
	if err != nil {
		t.Fatal(err)
	}
}
