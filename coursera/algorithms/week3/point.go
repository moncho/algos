package week4

import (
	"math"
)

type Point struct {
	X, Y int
}

func (p Point) slopeTo(that Point) float64 {
	if p.X == that.X {
		if p.Y == that.Y {
			return math.Inf(-1)
		}
		return math.Inf(1)
	} else if p.Y == that.Y {
		return 0
	}
	return float64(that.Y-p.Y) / float64(that.X-p.X)
}

/**
 * Compares two points by y-coordinate, breaking ties by X-coordinate. Formally, the invoking
 * point (X0, y0) is less than the argument point (X1, y1) if and only if either y0 < y1 or if
 * y0 = y1 and X0 < X1.
 *

 */
func (p Point) compareTo(that Point) int {
	if p.Y == that.Y {
		return p.X - that.X
	}
	return p.Y - that.Y

}

/**
 * Sorts the given points by the slope they make with this point.
 */
func (p Point) slopeOrder(pp []Point) func(i, j int) bool {
	return func(i, j int) bool {
		iSlope := p.slopeTo(pp[i])
		jSlope := p.slopeTo(pp[j])
		if iSlope == jSlope {
			return pp[i].compareTo(pp[j]) < 0
		}
		return iSlope < jSlope
	}
}
