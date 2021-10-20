package noise

import "image"

type ShuffleNoise struct {
	Intensity float64
	Image     image.Image
}

func (n ShuffleNoise) Run() image.Image {
	shuffled := image.NewGray(n.Image.Bounds())

	return shuffled
}

func CreateShuffleNoise(image image.Image, intensity float64) Noise {
	return &ShuffleNoise{
		Intensity: intensity,
		Image:     image,
	}
}
