package binarization

import (
	"image"
	"image/color"
)

// Otsu binarization implementation
type Otsu struct {
	Image *image.Gray
}

// CalculateHistogram creates histogram for stored image
func (o Otsu) CalculateHistogram() []uint8 {
	histogram := make([]uint8, 256)

	for y := o.Image.Bounds().Min.Y; y < o.Image.Bounds().Max.Y; y++ {
		for x := o.Image.Bounds().Min.X; x < o.Image.Bounds().Max.X; x++ {
			intensity := o.Image.GrayAt(x, y).Y
			histogram[intensity]++
		}
	}

	return histogram
}

// CalculateAverageIntensity calculates average intensity for stored image
func (o Otsu) CalculateAverageIntensity(histogram []uint8) float64 {
	sum := 0
	for y := o.Image.Bounds().Min.Y; y < o.Image.Bounds().Max.Y; y++ {
		for x := o.Image.Bounds().Min.X; x < o.Image.Bounds().Max.X; x++ {
			sum += int(o.Image.GrayAt(x, y).Y)
		}
	}
	return float64(sum) / float64(len(histogram))
}

// CalculateProbability calculates probability for image
func (o Otsu) CalculateProbability(histogram []uint8) []float64 {
	probability := make([]float64, 256)
	width, height := o.Image.Bounds().Dx(), o.Image.Bounds().Dy()
	s := width * height

	for i, value := range histogram {
		probability[i] = float64(value) / float64(s)
	}

	return probability
}

// CalculateThreshold calculates threshold for image binarization
func (o Otsu) CalculateThreshold() int {
	histogram := o.CalculateHistogram()
	probability := o.CalculateProbability(histogram)

	q1, q1next := probability[0], 0.0

	mu1, mu1next := 0.0, 0.0
	mu2, mu2next := 0.0, 0.0
	mu := o.CalculateAverageIntensity(histogram)

	betweenVariance := 0.0
	maxBetweenVariance := 0.0

	threshold := 0

	for i := 1; i < 255; i++ {
		q1next = q1 + probability[i+1]
		mu1next = (q1*mu1 + float64(i+1)*(probability[i+1])) / q1next
		mu2next = (mu - q1next*mu1next) / (1 - q1next)
		betweenVariance = q1 * (1 - q1) * ((mu1 - mu2) * (mu1 - mu2))

		if betweenVariance > maxBetweenVariance {
			maxBetweenVariance = betweenVariance
			threshold = i
		}

		q1 = q1next
		mu1 = mu1next
		mu2 = mu2next

		if q1next == 0 {
			mu1 = 0
		}
	}

	return threshold
}

// Process runs Otsu binarization algorithm
func (o Otsu) Process() *image.Gray {
	threshold := o.CalculateThreshold()
	resultImage := image.NewGray(o.Image.Bounds())

	for y := o.Image.Bounds().Min.Y; y < o.Image.Bounds().Max.Y; y++ {
		for x := o.Image.Bounds().Min.X; x < o.Image.Bounds().Max.X; x++ {
			intensity := o.Image.GrayAt(x, y).Y
			if intensity < uint8(threshold) {
				resultImage.Set(x, y, color.Gray{Y: 0})
			} else {
				resultImage.Set(x, y, color.Gray{Y: 255})
			}
		}
	}

	return resultImage
}

// CreateOtsuBinarization creates Otsu binarization exemplar
func CreateOtsuBinarization(img *image.Gray) Binarization {
	return &Otsu{Image: img}
}
