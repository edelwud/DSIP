package utils

import (
	"gonum.org/v1/gonum/mat"
	"io/ioutil"
	"path"
)

const (
	HighIntensity = 49
	LowIntensity  = 48
)

func ReadShape(filepath string) (*mat.VecDense, error) {
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

func WriteShape(vec *mat.VecDense, filepath string) error {
	arr := make([]byte, 0)
	for i := 0; i < vec.Len(); i++ {
		value := vec.AtVec(i)
		if value == 1 {
			arr = append(arr, HighIntensity)
		}
		if value == 0 {
			arr = append(arr, LowIntensity)
		}
		if (i+1)%6 == 0 {
			arr = append(arr, 10)
		}
	}

	err := ioutil.WriteFile(filepath, arr, 0777)
	if err != nil {
		return err
	}

	return nil
}

func ReadAllShapes(folder string) ([]*mat.VecDense, error) {
	dir, err := ioutil.ReadDir(folder)
	if err != nil {
		return nil, err
	}

	result := make([]*mat.VecDense, 0)
	for _, shapes := range dir {
		if shapes.IsDir() {
			continue
		}
		shape, err := ReadShape(path.Join(folder, shapes.Name()))
		if err != nil {
			return nil, err
		}
		result = append(result, shape)
	}

	return result, nil
}
