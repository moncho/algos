package week4

import (
	"sort"
)

func FastCollinearPoints(points []Point) []LineSegment {
	var segments []LineSegment

	for i := 0; i < len(points)-3; i++ {
		p := points[i]
		sort.Slice(points, sortBySlopeToPoint(p, points))
		count := 2
		minIndex, maxIndex := i, i
		slopeTo := p.slopeTo(points[1])
		for j := 2; j < len(points); j++ {
			if slopeTo == p.slopeTo(points[j]) {
				count++
				maxIndex = j
			} else {
				minIndex = j
				count = 2
			}
			slopeTo = p.slopeTo(points[j])
		}
		if count >= 4 {
			segments = append(segments, LineSegment{
				points[minIndex],
				points[maxIndex],
			})
			i = maxIndex
		}
	}
	return segments
}

func sortBySlopeToPoint(p Point, pp []Point) func(i, j int) bool {
	return func(i, j int) bool {

		iSlope := p.slopeTo(pp[i])
		jSlope := p.slopeTo(pp[j])
		if iSlope == jSlope {
			return pp[i].compareTo(pp[j]) < 0
		}
		return iSlope < jSlope
	}
}
