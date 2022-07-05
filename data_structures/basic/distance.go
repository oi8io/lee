package basic

import "math"

type Point struct {
	x float64
	y float64
}

func GetMinDistance(points []Point) (p1, p2 Point) {
	var theta float64
	for i := 0; i < len(points); i++ {
		for j := 0; j < len(points); j++ {
			if i == j { //同一个点不用对比
				continue
			}
			dist := Dist(points[i], points[j])
			if theta > dist {
				theta = dist
				p1 = points[i]
				p2 = points[j]
			}
		}
	}
	return
}

func Dist(p1, p2 Point) float64 {
	sqrt := math.Sqrt(math.Pow(p1.x-p2.x, 2) + math.Pow(p1.y-p2.y, 2))
	if sqrt < 0 {
		return 0 - sqrt
	}
	return sqrt
}
