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
		return 0
	} else if p.Y == that.Y {
		return math.Inf(1)
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
 * Compares two points by the slope they make with this point. The slope is defined as in the
 * slopeTo() method.
 *
 */
func slopeOrder(o1, o2 Point) int {
	slopeTo := o1.slopeTo(o2)
	if math.IsInf(slopeTo, 1) || math.IsInf(slopeTo, -1) {
		return 0
	} else if slopeTo > 0 {
		return 1
	} else if slopeTo < 0 {
		return -1
	}
	return 0
}
