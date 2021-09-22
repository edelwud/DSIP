package filters

import (
	"bytes"
	"github.com/h2non/bimg"
	"image"
	"image/color"
	"image/jpeg"
	"os"
)

func CreateHighImage(src string, dst string) ([]float64, error) {
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

	for y := 1; y < height-1; y++ {
		for x := 1; x < width-1; x++ {
			tl := color.GrayModel.Convert(source.At(x, y-1)).(color.Gray).Y
			tc := color.GrayModel.Convert(source.At(x, y-1)).(color.Gray).Y
			tr := color.GrayModel.Convert(source.At(x, y-1)).(color.Gray).Y
			cl := color.GrayModel.Convert(source.At(x-1, y)).(color.Gray).Y
			c := color.GrayModel.Convert(source.At(x, y)).(color.Gray).Y
			cr := color.GrayModel.Convert(source.At(x+1, y)).(color.Gray).Y
			bl := color.GrayModel.Convert(source.At(x, y+1)).(color.Gray).Y
			bc := color.GrayModel.Convert(source.At(x, y+1)).(color.Gray).Y
			br := color.GrayModel.Convert(source.At(x, y+1)).(color.Gray).Y

			final := c*9 - (tl + tc + tr + cl + cr + bl + bc + br)
			if final > 255 {
				final = 255
			}
			hist[final]++

			result.Set(x, y, color.RGBA{
				R: final,
				G: final,
				B: final,
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

	return hist, nil
}
