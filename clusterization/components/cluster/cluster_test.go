package cluster

import (
	"clusterization/components/binarization"
	"clusterization/components/blur"
	"clusterization/components/mask"
	"clusterization/components/utils"
	"testing"
)

const (
	ImageCluster              = "./../../resources/easy/P0001460.jpg"
	ImageClusterStore         = "./../../resources/easy_output/image_cluster.jpg"
	ImageClusterAnalyticStore = "./../../resources/easy_output/image_analytic_area.jpg"
)

func TestBasicCluster_CollectAnalyticCriteria(t *testing.T) {
	image, err := utils.ReadImageJpeg(ImageCluster)
	if err != nil {
		t.Fatal(err)
	}

	gaussianBlur := blur.CreateGaussianBlur(image, 5)
	grayImage := utils.GrayscaleImage(gaussianBlur.Process())
	bin := binarization.CreateThresholdBinarization(grayImage, 200)

	contourMask := mask.CreateContourAreaMask()
	connectedAreas := CreateConnectedAreasAnalyzer(image, grayImage, bin.Process(), contourMask, 1)
	figures := connectedAreas.FindConnectedAreas()

	cluster := CreateCluster(figures, 150)
	analyticArea := cluster.CollectAnalyticCriteria(CompactnessCriteria)
	analyticAreaImage := cluster.DrawAnalyticArea(analyticArea)
	err = utils.WriteImageJpeg(analyticAreaImage, ImageClusterAnalyticStore)
	if err != nil {
		t.Fatal(err)
	}

	clusters := cluster.FindClusters(analyticArea)
	coloredClusters := cluster.DrawClusters(clusters, image.Bounds().Dx(), image.Bounds().Dy())
	err = utils.WriteImageJpeg(coloredClusters, ImageClusterStore)
	if err != nil {
		t.Fatal(err)
	}

}
