package unionfind

import "testing"

func Test_weightedQuickUnion(t *testing.T) {

	qu := NewWeightedQuickUnion(10)
	qu.Union(1, 2)

	if !qu.Connected(1, 2) {
		t.Errorf("%d and %d should be connected", 1, 2)
	}
	if qu.Connected(1, 9) {
		t.Errorf("%d and %d should not be connected", 1, 9)
	}
	qu.Union(1, 2)
	qu.Union(4, 5)
	qu.Union(4, 2)
	qu.Union(7, 9)

	if !qu.Connected(1, 2) {
		t.Errorf("%d and %d should be connected", 1, 2)
	}
	if !qu.Connected(1, 4) {
		t.Errorf("%d and %d should be connected", 1, 4)
	}
	if !qu.Connected(1, 5) {
		t.Errorf("%d and %d should be connected", 1, 5)
	}
	if qu.Connected(1, 9) {
		t.Errorf("%d and %d should not be connected", 1, 9)
	}

}
