package utils

import (
	"gonum.org/v1/gonum/mat"
	"io/ioutil"
)

const (
	HighIntensity = 49
	LowIntensity  = 48
)

func ReadTrainingShape(filepath string) (*mat.VecDense, error) {
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	vector := make([]float64, 0)
	for _, intensity := range content {
		if intensity == HighIntensity {
			vector = append(vector, 1)
		}
		if intensity == LowIntensity {
			vector = append(vector, 0)
		}
	}

	return mat.NewVecDense(len(vector), vector), nil
}
