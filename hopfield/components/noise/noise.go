package noise

import "image"

type Noise interface {
	Run() *image.Gray
}
