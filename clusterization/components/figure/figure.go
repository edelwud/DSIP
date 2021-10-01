package figure

import "image"

type Figure struct {
	Snapshot *image.Gray
	Route    []image.Point
}

func (f Figure) Subtract() {

}
