package main

import (
	"filter/filters"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"log"
)

const ResourcesFolder = "resources"

func main() {
	negHist, err := filters.CreateNegativeImage(ResourcesFolder+"/image.jpg", ResourcesFolder+"/negative.jpg")
	if err != nil {
		fmt.Println(err)
		return
	}

	negValuer := make(plotter.ValueLabels, 256)
	for i, num := range negHist {
		negValuer[i].Value = num
		negValuer[i].Label = fmt.Sprintf("%4.4f", negValuer[i].Value)
	}

	negBox, err := plotter.NewHist(negValuer, 256)
	if err != nil {
		log.Panic(err)
	}

	highHist, err := filters.CreateHighImage(ResourcesFolder+"/image.jpg", ResourcesFolder+"/high.jpg")
	if err != nil {
		fmt.Println(err)
		return
	}

	highValuer := make(plotter.ValueLabels, 256)
	for i, num := range highHist {
		highValuer[i].Value = num
		highValuer[i].Label = fmt.Sprintf("%4.4f", highValuer[i].Value)
	}

	highBox, err := plotter.NewHist(highValuer, 256)
	if err != nil {
		log.Panic(err)
	}

	sobelHist, err := filters.CreateSobelImage(ResourcesFolder+"/image.jpg", ResourcesFolder+"/sobel.jpg")
	if err != nil {
		fmt.Println(err)
		return
	}

	sobelValuer := make(plotter.ValueLabels, 256)
	for i, num := range sobelHist {
		sobelValuer[i].Value = num
		sobelValuer[i].Label = fmt.Sprintf("%4.4f", sobelValuer[i].Value)
	}

	sobelBox, err := plotter.NewHist(sobelValuer, 256)
	if err != nil {
		log.Panic(err)
	}

	neg := plot.New()
	neg.Add(negBox)
	err = neg.Save(200, 200, ResourcesFolder+"/neg_hist.png")
	if err != nil {
		log.Panic(err)
	}

	highPlot := plot.New()
	highPlot.Add(highBox)
	err = highPlot.Save(200, 200, ResourcesFolder+"/high_hist.png")
	if err != nil {
		log.Panic(err)
	}

	sobelPlot := plot.New()
	sobelPlot.Add(sobelBox)
	err = sobelPlot.Save(200, 200, ResourcesFolder+"/sobel_hist.png")
	if err != nil {
		log.Panic(err)
	}

	source := canvas.NewImageFromFile(ResourcesFolder + "/image.jpg")
	source.FillMode = canvas.ImageFillStretch
	source.SetMinSize(fyne.Size{
		Width:  200,
		Height: 200,
	})

	high := canvas.NewImageFromFile(ResourcesFolder + "/high.jpg")
	high.FillMode = canvas.ImageFillStretch
	high.SetMinSize(fyne.Size{
		Width:  200,
		Height: 200,
	})

	highCanvas := canvas.NewImageFromFile(ResourcesFolder + "/high_hist.png")
	highCanvas.FillMode = canvas.ImageFillStretch
	highCanvas.SetMinSize(fyne.Size{
		Width:  200,
		Height: 200,
	})

	negative := canvas.NewImageFromFile(ResourcesFolder + "/negative.jpg")
	negative.FillMode = canvas.ImageFillStretch
	negative.SetMinSize(fyne.Size{
		Width:  200,
		Height: 200,
	})

	negCanvas := canvas.NewImageFromFile(ResourcesFolder + "/neg_hist.png")
	negCanvas.FillMode = canvas.ImageFillStretch
	negCanvas.SetMinSize(fyne.Size{
		Width:  200,
		Height: 200,
	})

	sobel := canvas.NewImageFromFile(ResourcesFolder + "/sobel.jpg")
	sobel.FillMode = canvas.ImageFillStretch
	sobel.SetMinSize(fyne.Size{
		Width:  200,
		Height: 200,
	})

	sobelCanvas := canvas.NewImageFromFile(ResourcesFolder + "/sobel_hist.png")
	sobelCanvas.FillMode = canvas.ImageFillStretch
	sobelCanvas.SetMinSize(fyne.Size{
		Width:  200,
		Height: 200,
	})

	a := app.New()
	w := a.NewWindow("Image filtering")
	w.SetContent(container.NewVBox(
		source,
		container.NewAppTabs(
			container.NewTabItem("High filter", container.NewVBox(
				high,
				highCanvas,
			)),
			container.NewTabItem("Negative filter", container.NewVBox(
				negative,
				negCanvas,
			)),
			container.NewTabItem("Sobel filter", container.NewVBox(
				sobel,
				sobelCanvas,
			)),
		),
	))
	w.Resize(fyne.Size{
		Width:  500,
		Height: 450,
	})
	w.ShowAndRun()
}
