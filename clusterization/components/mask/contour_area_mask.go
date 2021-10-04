package mask

import "image"

// ContourAreaMask implements counter area mask for connected area analysis
type ContourAreaMask struct {
	Mask []image.Point
}

// Generate generates mask for contour area analysis
func (m ContourAreaMask) Generate(scale int) []image.Point {
	m.Mask = make([]image.Point, 0)
	m.Mask = append(m.Mask, image.Point{X: 0, Y: 0})

	x := 0
	y := -1

	var pawn bool
	width := 1
	height := 1

	for width <= scale {
		xFlag := false
		yFlag := true
		flag := -width
		x++
		y++

		m.Mask = append(m.Mask, image.Point{X: x, Y: 0})

		for x != width || y != -1 {
			if flag > 0 {
				if x == width {
					xFlag = false
				}
				if x == -width {
					xFlag = true
				}
				if xFlag {
					x++
				} else {
					x--
				}
				pawn = true
				flag--
			} else {
				if y == height {
					yFlag = false
				}
				if y == -height {
					yFlag = true
				}
				if yFlag {
					y++
				} else {
					y--
				}
				pawn = false
				flag++
			}
			m.Mask = append(m.Mask, image.Point{X: x, Y: y})
			if flag == 0 {
				if pawn {
					flag = -2 * width
				} else {
					flag = +2 * width
				}
			}
		}
		width++
		height++
	}

	return m.Mask
}

// CreateContourAreaMask initialized ContourAreaMask exemplar
func CreateContourAreaMask() AreaMask {
	return &ContourAreaMask{
		Mask: nil,
	}
}
