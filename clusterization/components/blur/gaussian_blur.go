package blur

import (
	"image"
	"image/color"
	"math"
)

// GaussianBlur declares Gaussian blur implementation
type GaussianBlur struct {
	Image  image.Image
	Radius int
}

// GenerateKernel generates Gaussian kernel
func (b GaussianBlur) GenerateKernel() [][]float64 {
	kernelWidth := 2*b.Radius + 1

	kernel := make([][]float64, kernelWidth)
	for i := range kernel {
		kernel[i] = make([]float64, kernelWidth)
	}

	sigma := math.Max(float64(b.Radius)/2, 1)
	sum := 0.0

	for x := -b.Radius; x < b.Radius; x++ {
		for y := -b.Radius; y < b.Radius; y++ {
			exponentNumerator := -(x*x + y*y)
			exponentDenominator := 2 * sigma * sigma
			exponentExpression := math.Exp(float64(exponentNumerator) / exponentDenominator)
			kernelValue := exponentExpression / (2 * math.Pi * sigma * sigma)
			kernel[x+b.Radius][y+b.Radius] = kernelValue
			sum += kernelValue
		}
	}

	for i, row := range kernel {
		for j := range row {
			kernel[i][j] /= sum
		}
	}

	return kernel
}

// Process precesses Gaussian blur
func (b GaussianBlur) Process() image.Image {
	blurred := image.NewRGBA(b.Image.Bounds())
	kernel := b.GenerateKernel()

	for x := blurred.Bounds().Min.X + b.Radius; x < blurred.Bounds().Max.X-b.Radius; x++ {
		for y := blurred.Bounds().Min.Y + b.Radius; y < blurred.Bounds().Max.Y-b.Radius; y++ {
			redSummary := 0.0
			greenSummary := 0.0
			blueSummary := 0.0
			alphaSummary := 0.0

			for kernelX := -b.Radius; kernelX < b.Radius; kernelX++ {
				for kernelY := -b.Radius; kernelY < b.Radius; kernelY++ {
					kernelValue := kernel[kernelX+b.Radius][kernelY+b.Radius]
					red, green, blue, a := b.Image.At(x-kernelX, y-kernelY).RGBA()
					red, green, blue, a = red>>8, green>>8, blue>>8, a>>8

					redSummary += float64(red) * kernelValue
					greenSummary += float64(green) * kernelValue
					blueSummary += float64(blue) * kernelValue
					alphaSummary += float64(a) * kernelValue
				}
			}

			blurred.Set(x, y, color.RGBA{
				R: uint8(redSummary),
				G: uint8(greenSummary),
				B: uint8(blueSummary),
				A: uint8(alphaSummary),
			})
		}
	}

	return blurred
}

// CreateGaussianBlur initializes GaussianBlur exemplar
func CreateGaussianBlur(img image.Image, radius int) Blur {
	return &GaussianBlur{
		Image:  img,
		Radius: radius,
	}
}
