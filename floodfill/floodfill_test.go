package main

import (
	"math/rand"
	"testing"
	"time"
)

func TestMaxConsecutiveValues(t *testing.T) {

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
				t.Errorf("maxConsecutiveValues() count got = %v, want %v", gotCount, tt.expected.count)
			}
		})
	}
}

func BenchmarkMaxConsecutiveValues(b *testing.B) {

	benchmarks := []struct {
		name          string
		rows, columns int
	}{
		{"10x5", 10, 5},
		{"100x100", 100, 100},
		{"1000x1000", 1000, 1000},
	}
	rand.Seed(time.Now().UnixNano())
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			b.StopTimer()
			grid := randomGrid(bm.rows, bm.columns, 3)
			b.StartTimer()
			for i := 0; i < b.N; i++ {
				maxConnectedValue(grid)
			}
		})
	}

}
