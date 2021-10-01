package cluster

import "image"

// AreaMask declares basic interface for mask generation algorithms
type AreaMask interface {
	Generate(scale int) []image.Point
}
