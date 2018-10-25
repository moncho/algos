package week4

import (
	"sort"
)

func FastCollinearPoints(points []Point) []LineSegment {
	var segments []LineSegment
	sorted := make([]Point, len(points))
	copy(sorted, points)
	for i := 0; i < len(points); i++ {
		p := points[i]
		sort.Slice(sorted, sortBySlopeToPoint(p, sorted))
		minIndex, maxIndex := 1, 1
		slopeTo := p.slopeTo(sorted[1])
		for j := 2; j < len(points); j++ {
			if slopeTo == p.slopeTo(sorted[j]) {
				maxIndex = j
			} else {
				if maxIndex-minIndex >= 2 && p.compareTo(sorted[minIndex]) < 0 {
					segments = append(segments, LineSegment{
						p,
						sorted[maxIndex],
					})
				}
				minIndex = j
				slopeTo = p.slopeTo(sorted[j])
			}
		}
		if maxIndex-minIndex >= 2 && p.compareTo(sorted[minIndex]) < 0 {
			segments = append(segments, LineSegment{
				p,
				sorted[maxIndex],
			})
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
