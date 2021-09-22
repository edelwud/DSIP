package filters

import (
	"bytes"
	"github.com/h2non/bimg"
	"image"
	"image/color"
	"image/jpeg"
	"os"
)

func CreateNegativeImage(src string, dst string) ([]float64, error) {
	hist := make([]float64, 256)

	buffer, err := bimg.Read(src)
	if err != nil {
		return nil, err
	}

	source, _, err := image.Decode(bytes.NewReader(buffer))
	if err != nil {
		return nil, err
	}

	bounds := source.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	result := image.NewRGBA(bounds)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			pixel := color.GrayModel.Convert(source.At(x, y)).(color.Gray)
			negative := 255 - pixel.Y
			hist[negative]++
			result.Set(x, y, color.RGBA{
				R: negative,
				G: negative,
				B: negative,
				A: 1,
			})
		}
	}

	file, err := os.Create(dst)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)

	err = jpeg.Encode(file, result, nil)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
