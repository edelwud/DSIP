package utils

import (
	"image/color"
	"math/rand"
)

func GenerateRandomColour() color.Color {
	return color.RGBA{
		R: uint8(rand.Int() % 255),
		G: uint8(rand.Int() % 255),
		B: uint8(rand.Int() % 255),
		A: 1,
	}
}
