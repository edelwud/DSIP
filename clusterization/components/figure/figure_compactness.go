package figure

// Compactness declares minimum interface for compactness functionality implementation
type Compactness interface {
	FindSquare() int
	FindPerimeter() int
	FindCompactness() int
}

// FindSquare finds figure square
func (f Figure) FindSquare() int {
	return len(f.Relative)
}

// FindPerimeter finds figure perimeter
func (f Figure) FindPerimeter() int {
	perimeterRoute := make(map[int][]int, 0)
	prevIntensity := 0

	for y := f.Snapshot.Bounds().Min.Y - 1; y < f.Snapshot.Bounds().Max.Y+1; y++ {
		for x := f.Snapshot.Bounds().Min.X - 1; x < f.Snapshot.Bounds().Max.X+1; x++ {
			currIntensity := int(f.Snapshot.GrayAt(x, y).Y)
			if currIntensity != prevIntensity {
				perimeterRoute[x] = append(perimeterRoute[x], y)
			}
			prevIntensity = currIntensity
		}
		prevIntensity = 0
	}

	for x := f.Snapshot.Bounds().Min.X - 1; x < f.Snapshot.Bounds().Max.X+1; x++ {
		for y := f.Snapshot.Bounds().Min.Y - 1; y < f.Snapshot.Bounds().Max.Y+1; y++ {
			currIntensity := int(f.Snapshot.GrayAt(x, y).Y)
			if currIntensity != prevIntensity {
				found := false
				for _, y1 := range perimeterRoute[x] {
					if y1 == y {
						found = true
					}
				}
				if !found {
					perimeterRoute[x] = append(perimeterRoute[x], y)
				}
			}
			prevIntensity = currIntensity
		}
		prevIntensity = 0
	}

	perimeter := 0
	for _, ySlice := range perimeterRoute {
		perimeter += len(ySlice)
	}

	return perimeter
}

// FindCompactness finds figure compactness
func (f Figure) FindCompactness() int {
	perimeter := f.FindPerimeter()
	square := f.FindSquare()
	return perimeter * perimeter / square
}
