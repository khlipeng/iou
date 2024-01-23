package iou

import (
	"math"
)

type Point [2]float64

type Polygon []Point

func (p Polygon) Area() float64 {
	area := 0.0
	n := len(p)
	for i := 0; i < n; i++ {
		j := (i + 1) % n
		area += p[i][0] * p[j][1]
		area -= p[j][0] * p[i][1]
	}
	area = math.Abs(area) / 2.0
	return area
}

func (p Polygon) Intersection(poly2 Polygon) float64 {
	var minX, minY, maxX, maxY float64
	// 计算最小、最大坐标
	for _, point := range append(p, poly2...) {
		minX = math.Min(minX, point[0])
		minY = math.Min(minY, point[1])
		maxX = math.Max(maxX, point[0])
		maxY = math.Max(maxY, point[1])
	}

	// 遍历矩形区域内的每个像素点，并判断是否在交集内
	intersectArea := 0.0
	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			if p.In(x, y) && poly2.In(x, y) {
				intersectArea++
			}
		}
	}

	// 计算交集面积
	pixelArea := 1.0
	intersectArea *= pixelArea
	return intersectArea
}

func (p Polygon) In(x float64, y float64) bool {
	inside := false
	n := len(p)
	j := n - 1
	for i := 0; i < n; i++ {
		if ((p[i][1] > y) != (p[j][1] > y)) && (x < (p[j][0]-p[i][0])*(y-p[i][1])/(p[j][1]-p[i][1])+p[i][0]) {
			inside = !inside
		}
		j = i
	}

	return inside
}

func (p *Polygon) UnionByIntersection(other Polygon, intersection float64) float64 {
	return p.Area() + other.Area() - intersection
}

func (p *Polygon) Union(other Polygon) float64 {
	return p.Area() + other.Area() - p.Intersection(other)
}

func (p *Polygon) IoU(other Polygon) float64 {
	intersection := p.Intersection(other)
	union := p.UnionByIntersection(other, intersection)
	return intersection / union
}
