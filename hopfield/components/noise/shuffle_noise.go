package noise

import (
	"image"
	"image/color"
	"math/rand"
)

type ShuffleNoise struct {
	Intensity float64
	Image     *image.Gray
}

func (n ShuffleNoise) Run() image.Image {
	shuffled := image.NewGray(n.Image.Bounds())

	x, y := shuffled.Bounds().Dx(), shuffled.Bounds().Dy()

	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			c := n.Image.GrayAt(i, j).Y
			pattern := n.Intensity - rand.Float64()
			if pattern >= 0 {
				shuffled.Set(i, j, color.Gray{Y: 255 - c})
			} else {
				shuffled.Set(i, j, color.Gray{Y: c})
			}
		}
	}

	return shuffled
}

func CreateShuffleNoise(image *image.Gray, intensity float64) Noise {
	return &ShuffleNoise{
		Intensity: intensity,
		Image:     image,
	}
}
