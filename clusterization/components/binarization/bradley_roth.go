package binarization

import "image"

// BradleyRoth algorithm implementation
type BradleyRoth struct {
	Image image.Image
}

// Process Bradley-Roth binarization
func (br BradleyRoth) Process() (image.Image, error) {
	return nil, nil
}

// CreateBradleyRothBinarization creates BradleyRoth exemplar
func CreateBradleyRothBinarization(img image.Image) Binarization {
	return &BradleyRoth{Image: img}
}
