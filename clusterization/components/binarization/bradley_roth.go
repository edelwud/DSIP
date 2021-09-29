package binarization

import (
	"image"
	"image/color"
)

// BradleyRoth algorithm implementation
type BradleyRoth struct {
	Image *image.Gray
	T     float64
}

func (br BradleyRoth) CreateIntegralImage() *image.Gray {
	integralImg := image.NewGray(br.Image.Bounds())

	for y := br.Image.Bounds().Min.Y; y < br.Image.Bounds().Max.Y; y++ {
		sum := uint8(0)
		for x := br.Image.Bounds().Min.X; x < br.Image.Bounds().Max.X; x++ {
			sum += br.Image.GrayAt(x, y).Y
			if x == 0 {
				integralImg.Set(x, y, color.Gray{Y: sum})
			} else {
				prevIntensity := br.Image.GrayAt(x-1, y).Y
				integralImg.Set(x, y, color.Gray{Y: sum + prevIntensity})
			}
		}
	}

	return integralImg
}

// Process Bradley-Roth binarization
func (br BradleyRoth) Process() *image.Gray {
	integralImg := br.CreateIntegralImage()
	resultImg := image.NewGray(integralImg.Bounds())

	rect := br.Image.Bounds()
	width, height := rect.Size().X, rect.Size().Y
	s := width / 8
	s2 := s / 2

	sum := uint8(0)
	count := uint8(0)
	x1, y1, x2, y2 := 0, 0, 0, 0

	for y := br.Image.Bounds().Min.Y; y < br.Image.Bounds().Max.Y; y++ {
		y1 = y - s2
		y2 = y + s2
		if y1 < 0 {
			y1 = 0
		}
		if y2 >= height {
			y2 = height - 1
		}
		for x := br.Image.Bounds().Min.X; x < br.Image.Bounds().Max.X; x++ {
			x1 = x - s2
			x2 = x + s2
			if x1 < 0 {
				x1 = 0
			}
			if x2 >= width {
				x2 = width - 1
			}

			count = uint8((x2 - x1) * (y2 - y1))
			topRight := integralImg.GrayAt(x2, y1).Y
			topLeft := integralImg.GrayAt(x1, y1).Y
			bottomRight := integralImg.GrayAt(x2, y2).Y
			bottomLeft := integralImg.GrayAt(x1, y2).Y
			sum = bottomRight - topRight - bottomLeft + topLeft

			if float64(br.Image.GrayAt(x, y).Y*count) < float64(sum)*(1.0-br.T) {
				resultImg.Set(x, y, color.Gray{Y: 0})
			} else {
				resultImg.Set(x, y, color.Gray{Y: 255})
			}
		}
	}

	return resultImg
}

// CreateBradleyRothBinarization creates BradleyRoth exemplar
func CreateBradleyRothBinarization(img *image.Gray, t float64) Binarization {
	return &BradleyRoth{Image: img, T: t}
}
