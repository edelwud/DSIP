package utils

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"testing"
)

func TestReadTrainingShape(t *testing.T) {
	shape, err := ReadShape("../../resources/training/1.txt")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(mat.Formatted(shape.T()))
}

func TestWriteShape(t *testing.T) {
	shape, err := ReadShape("../../resources/training/1.txt")
	if err != nil {
		t.Fatal(err)
	}

	err = WriteShape(shape, "../../resources/results/1.txt")
	if err != nil {
		t.Fatal(err)
	}
}
