package cluster

import (
	"clusterization/components/binarization"
	"clusterization/components/utils"
	"testing"
)

const (
	ImageArrangeShapes      = "./../../resources/easy/P0001460.jpg"
	ImageArrangeShapesStore = "./../../resources/easy_output/image_arranged_shapes.jpg"
)

func TestArrangeShapes(t *testing.T) {
	image, err := utils.ReadImageJpeg(ImageArrangeShapes)
	if err != nil {
		t.Fatal(err)
	}

	grayImage := utils.GrayscaleImage(image)
	bin := binarization.CreateThresholdBinarization(grayImage, 200)

	connectedAreas := CreateConnectedAreasAnalyzer(bin.Process(), 30)
	figures := connectedAreas.FindConnectedAreas()

	x, y := connectedAreas.Image.Bounds().Max.X, connectedAreas.Image.Bounds().Max.Y

	arrangedShapes := ArrangeShapes(figures, x, y)
	err = utils.WriteImageJpeg(arrangedShapes, ImageArrangeShapesStore)
	if err != nil {
		t.Fatal(err)
	}
}
