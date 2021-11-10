package kmeans

import (
	"math"
	"math/rand"
)

type Point struct {
	Entry []float64
}

type Centroid struct {
	Center Point
	Points []Point
}

func (p1 Point) distanceTo(p2 Point) float64 {
	sum := float64(0)
	for e := 0; e < len(p1.Entry); e++ {
		sum += math.Pow(p1.Entry[e]-p2.Entry[e], 2)
	}
	return math.Sqrt(sum)
}

func (c1 *Centroid) reCenter() float64 {
	newCentroid := make([]float64, len(c1.Center.Entry))
	for _, e := range c1.Points {
		for r := 0; r < len(newCentroid); r++ {
			newCentroid[r] += e.Entry[r]
		}
	}
	for r := 0; r < len(newCentroid); r++ {
		newCentroid[r] /= float64(len(c1.Points))
	}
	oldCenter := c1.Center
	c1.Center = Point{newCentroid}
	return oldCenter.distanceTo(c1.Center)
}

func KMeans(data []Point, k uint64, DELTA float64) (Centroids []Centroid) {
	rand.Seed(1)
	for i := uint64(0); i < k; i++ {
		Centroids = append(Centroids, Centroid{Center: data[rand.Intn(len(data))]})
	}

	converged := false
	for !converged {
		for i := range data {
			minDistance := math.MaxFloat64
			z := 0
			for v, e := range Centroids {
				distance := data[i].distanceTo(e.Center)
				if distance < minDistance {
					minDistance = distance
					z = v
				}
			}
			Centroids[z].Points = append(Centroids[z].Points, data[i])
		}
		maxDelta := -math.MaxFloat64
		for i := range Centroids {
			movement := Centroids[i].reCenter()
			if movement > maxDelta {
				maxDelta = movement
			}
		}
		if DELTA >= maxDelta {
			converged = true
			return
		}
		for i := range Centroids {
			Centroids[i].Points = nil
		}
	}
	return Centroids
}
