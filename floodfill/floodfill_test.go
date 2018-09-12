package main

import "testing"

func Test_maxConsecutiveValues(t *testing.T) {

	type expected struct {
		value int
		count int
	}
	tests := []struct {
		name     string
		grid     [][]int
		expected expected
	}{
		{
			"1x1",
			[][]int{
				{0},
			},
			expected{0, 1},
		},
		{
			"1x5",
			[][]int{
				{0, 1, 2, 1, 1},
			},
			expected{1, 2},
		},
		{
			"3x5",
			[][]int{
				{0, 0, 1, 1, 2},
				{0, 0, 0, 0, 0},
				{0, 0, 1, 1, 2},
			},
			expected{0, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValue, gotCount := maxConnectedValue(tt.grid)
			if gotValue != tt.expected.value {
				t.Errorf("maxConsecutiveValues() value got = %v, want %v", gotValue, tt.expected.value)
			}
			if gotCount != tt.expected.count {
				t.Errorf("maxConsecutiveValues() count got1 = %v, want %v", gotCount, tt.expected.count)
			}
		})
	}
}

func BenchmarkMaxConsecutiveValues10x5(b *testing.B) {
	grid := randomGrid(10, 5, 3)
	var value, max int
	for n := 0; n < b.N; n++ {
		value, max = maxConnectedValue(grid)
	}
	_ = value
	_ = max
}
func BenchmarkMaxConsecutiveValues100x100(b *testing.B) {
	grid := randomGrid(100, 100, 3)
	var value, max int
	for n := 0; n < b.N; n++ {
		value, max = maxConnectedValue(grid)
	}
	_ = value
	_ = max
}

func BenchmarkMaxConsecutiveValues1000x1000(b *testing.B) {
	grid := randomGrid(1000, 1000, 3)
	var value, max int
	for n := 0; n < b.N; n++ {
		value, max = maxConnectedValue(grid)
	}
	_ = value
	_ = max
}
