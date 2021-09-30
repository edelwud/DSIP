package cluster

import (
	binarization2 "clusterization/components/binarization"
	"clusterization/components/utils"
	"testing"
)

const (
	ImageConnectedAreas      = "./../../resources/easy/P0001460.jpg"
	ImageConnectedAreasStore = "./../../resources/easy_output/image_connected_areas.jpg"
)

func TestConnectedAreas_FindConnectedAreas(t *testing.T) {
	image, err := utils.ReadImageJpeg(ImageConnectedAreas)
	if err != nil {
		t.Fatal(err)
	}

	grayImage := utils.GrayscaleImage(image)
	binarization := binarization2.CreateOtsuBinarization(grayImage)

	connectedAreas := CreateConnectedAreasAnalyzer(binarization.Process())
	connectedAreas.FindConnectedAreas()

	err = utils.WriteImageJpeg(connectedAreas.Area, ImageConnectedAreasStore)
	if err != nil {
		t.Fatal(err)
	}
}
