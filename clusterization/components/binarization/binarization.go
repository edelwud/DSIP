package binarization

import "image"

// Binarization declares basic interface for binarization algorithms
type Binarization interface {
	Process() (*image.Gray, error)
}
