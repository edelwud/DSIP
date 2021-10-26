package main

import (
	"hopfield/components/binarization"
	"hopfield/components/hopfield"
	"hopfield/components/noise"
	"hopfield/components/utils"
	"image"
	"io/ioutil"
	"strconv"
	"testing"
)

func TestSyncHopfield(t *testing.T) {
	trainingDir, err := ioutil.ReadDir("./resources/training")
	if err != nil {
		t.Fatal(err)
	}

	goldenObjects := make([]*image.Gray, 0)
	for _, file := range trainingDir {
		png, err := utils.ReadImagePNG("./resources/training/" + file.Name())
		if err != nil {
			t.Fatal(err)
		}

		gray := utils.GrayscaleImage(png)
		binRunner := binarization.CreateThresholdBinarization(gray, 128)
		goldenObjects = append(goldenObjects, binRunner.Process())
	}

	for index, img := range goldenObjects {
		for i := 0.0; i < 1.0; i += 0.1 {
			noiseRunner := noise.CreateShuffleNoise(img, i)
			percent := int(i * 100)
			err := utils.WriteImagePNG(
				noiseRunner.Run(),
				"./resources/shuffle/train_"+strconv.Itoa(index+1)+"_"+strconv.Itoa(percent)+".png")
			if err != nil {
				t.Fatal(err)
			}
		}
	}

	shuffleDir, err := ioutil.ReadDir("./resources/shuffle")
	if err != nil {
		t.Fatal(err)
	}

	for _, file := range shuffleDir {
		if file.Name() == ".gitkeep" {
			continue
		}

		png, err := utils.ReadImagePNG("./resources/shuffle/" + file.Name())
		if err != nil {
			t.Fatal(err)
		}

		gray := utils.GrayscaleImage(png)
		binRunner := binarization.CreateThresholdBinarization(gray, 128)

		result := hopfield.SyncHopfield(binRunner.Process(), goldenObjects)
		err = utils.WriteImagePNG(result, "./resources/results/sync_"+file.Name())
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestAsyncHopfield(t *testing.T) {
	trainingDir, err := ioutil.ReadDir("./resources/training")
	if err != nil {
		t.Fatal(err)
	}

	goldenObjects := make([]*image.Gray, 0)
	for _, file := range trainingDir {
		png, err := utils.ReadImagePNG("./resources/training/" + file.Name())
		if err != nil {
			t.Fatal(err)
		}

		gray := utils.GrayscaleImage(png)
		binRunner := binarization.CreateThresholdBinarization(gray, 128)
		goldenObjects = append(goldenObjects, binRunner.Process())
	}

	for index, img := range goldenObjects {
		for i := 0.0; i < 1.0; i += 0.1 {
			noiseRunner := noise.CreateShuffleNoise(img, i)
			percent := int(i * 100)
			err := utils.WriteImagePNG(
				noiseRunner.Run(),
				"./resources/shuffle/train_"+strconv.Itoa(index+1)+"_"+strconv.Itoa(percent)+".png")
			if err != nil {
				t.Fatal(err)
			}
		}
	}

	shuffleDir, err := ioutil.ReadDir("./resources/shuffle")
	if err != nil {
		t.Fatal(err)
	}

	for _, file := range shuffleDir {
		if file.Name() == ".gitkeep" {
			continue
		}

		png, err := utils.ReadImagePNG("./resources/shuffle/" + file.Name())
		if err != nil {
			t.Fatal(err)
		}

		gray := utils.GrayscaleImage(png)
		binRunner := binarization.CreateThresholdBinarization(gray, 128)

		result := hopfield.AsyncHopfield(binRunner.Process(), goldenObjects)
		err = utils.WriteImagePNG(result, "./resources/results/async_"+file.Name())
		if err != nil {
			t.Fatal(err)
		}
	}
}
