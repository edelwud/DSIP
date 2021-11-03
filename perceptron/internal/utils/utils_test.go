package utils

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"testing"
)

func TestReadTrainingShape(t *testing.T) {
	shape, err := ReadTrainingShape("../../resources/training/1.txt")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(mat.Formatted(shape.T()))
}
