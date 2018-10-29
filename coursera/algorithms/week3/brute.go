package week4

import (
	"math"
)

type LineSegment struct {
	P, Q Point
}

// BruteCollinearPoints finds all line segments containing 4 points
func BruteCollinearPoints(points []Point) []LineSegment {
	var segments []LineSegment
	var slopes []float64
	pLen := len(points)

	for p := 0; p < pLen-3; p++ {
		for q := p + 1; q < pLen-2; q++ {
			collinear := false
			slopeTo := points[p].slopeTo(points[q])
			minPoint, maxPoint := sortPoints(points[p], points[q])
			if !isDegenerate(slopeTo) {
				for r := q + 1; r < pLen-1; r++ {
					if slopeTo == points[q].slopeTo(points[r]) {
						loopMin, loopMax := sortPoints(points[q], points[r])
						minPoint, _ = sortPoints(loopMin, minPoint)
						_, maxPoint = sortPoints(loopMax, maxPoint)
						for s := r + 1; s < pLen; s++ {
							if slopeTo == points[r].slopeTo(points[s]) {
								collinear = true
								loopMin, loopMax = sortPoints(points[r], points[s])
								minPoint, _ = sortPoints(loopMin, minPoint)
								_, maxPoint = sortPoints(loopMax, maxPoint)
							}
						}
					}
				}
			}
			if collinear && isNew(slopes, slopeTo) {
				slopes = append(slopes, slopeTo)
				segments = append(segments, LineSegment{minPoint,
					maxPoint})
			}

		}
	}
	return segments
}

func isNew(slopes []float64, slope float64) bool {
	for _, s := range slopes {
		if s == slope {
			return false
		}
	}
	return true
}

func sortPoints(p, q Point) (Point, Point) {
	if p.compareTo(q) <= 0 {
		return p, q
	}
	return q, p
}

func isDegenerate(f float64) bool {
	return math.IsInf(f, -1)

}
