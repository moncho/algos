package week4

import (
	"math"
)

type LineSegment struct {
	P, Q Point
}

// finds all line segments containing 4 points
func BruteCollinearPoints(points []Point) []LineSegment {
	var segments []LineSegment
	pLen := len(points)

	for p := 0; p < pLen-3; p++ {
		for q := p + 1; q < pLen-2; q++ {
			slopeTo := points[p].slopeTo(points[q])
			if !isDegenerate(slopeTo) {
				var minPoint, maxPoint Point
				if points[p].compareTo(points[q]) <= 0 {
					minPoint = points[p]
					maxPoint = points[q]
				} else {
					minPoint = points[q]
					maxPoint = points[p]
				}

				for r := q + 1; r < pLen-1; r++ {
					if slopeTo == points[q].slopeTo(points[r]) {
						for s := r + 1; s < pLen; s++ {
							if slopeTo == points[r].slopeTo(points[s]) {
								segments = append(segments, LineSegment{minPoint,
									maxPoint})
								break
							}
						}
					}
				}
			}
		}
	}
	return segments
}

func isDegenerate(f float64) bool {
	return math.IsInf(f, -1)

}
