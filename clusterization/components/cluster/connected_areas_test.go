package cluster

import (
	"clusterization/components/binarization"
	"clusterization/components/blur"
	"clusterization/components/mask"
	"clusterization/components/utils"
	"strconv"
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

	gaussianBlur := blur.CreateGaussianBlur(image, 5)
	grayImage := utils.GrayscaleImage(gaussianBlur.Process())
	bin := binarization.CreateThresholdBinarization(grayImage, 200)

	contourMask := mask.CreateContourAreaMask()
	connectedAreas := CreateConnectedAreasAnalyzer(image, grayImage, bin.Process(), contourMask, 1)
	figures := connectedAreas.FindConnectedAreas()

	err = utils.WriteImageJpeg(connectedAreas.Image, ImageConnectedAreasStore)
	if err != nil {
		t.Fatal(err)
	}

	for i, figure := range figures {
		err = utils.WriteImageJpeg(figure.Snapshot, "./../../resources/easy_output/image_connected_areas_"+strconv.Itoa(i)+".jpg")
		if err != nil {
			t.Fatal(err)
		}
	}
}
