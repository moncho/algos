package main

import "testing"

func Test_weightedQuickUnion(t *testing.T) {

	qu := NewWeightedQuickUnion(10)
	qu.union(1, 2)

	if !qu.connected(1, 2) {
		t.Errorf("%d and %d should be connected", 1, 2)
	}
	if qu.connected(1, 9) {
		t.Errorf("%d and %d should not be connected", 1, 9)
	}
	qu.union(1, 2)
	qu.union(4, 5)
	qu.union(4, 2)
	qu.union(7, 9)

	if !qu.connected(1, 2) {
		t.Errorf("%d and %d should be connected", 1, 2)
	}
	if !qu.connected(1, 4) {
		t.Errorf("%d and %d should be connected", 1, 4)
	}
	if qu.connected(1, 5) {
		t.Errorf("%d and %d should be connected", 1, 5)
	}
	if qu.connected(1, 9) {
		t.Errorf("%d and %d should not be connected", 1, 9)
	}

}
