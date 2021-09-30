package cluster

import (
	"clusterization/components/binarization"
	"clusterization/components/utils"
	"strconv"
	"testing"
)

const (
	ImageConnectedAreas      = "./../../resources/easy/P0001468.jpg"
	ImageConnectedAreasStore = "./../../resources/easy_output/image_connected_areas.jpg"
)

func TestConnectedAreas_FindConnectedAreas(t *testing.T) {
	image, err := utils.ReadImageJpeg(ImageConnectedAreas)
	if err != nil {
		t.Fatal(err)
	}

	grayImage := utils.GrayscaleImage(image)
	bin := binarization.CreateOtsuBinarization(grayImage)

	connectedAreas := CreateConnectedAreasAnalyzer(bin.Process(), 30)
	connectedAreas.FindConnectedAreas()

	err = utils.WriteImageJpeg(connectedAreas.Image, ImageConnectedAreasStore)
	if err != nil {
		t.Fatal(err)
	}

	for i, figure := range connectedAreas.Figures {
		err = utils.WriteImageJpeg(figure.Snapshot, "./../../resources/easy_output/image_connected_areas_"+strconv.Itoa(i)+".jpg")
		if err != nil {
			t.Fatal(err)
		}
	}
}
