package blur

import "image"

// Blur declares basic interface for blur algorithms
type Blur interface {
	Process() image.Image
}
