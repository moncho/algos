package main

import (
	"testing"
)

var (
	m = 50000
	n = 100000
)

func BenchmarkSample(b *testing.B) {
	var res []int
	for i := 0; i < b.N; i++ {
		res = Sample(m, n)
	}
	_ = res
}

func BenchmarkSampleF2(b *testing.B) {
	var res []int
	for i := 0; i < b.N; i++ {
		res = SampleF2(m, n)
	}
	_ = res
}

func BenchmarkSampleF2Map(b *testing.B) {
	var res []int
	for i := 0; i < b.N; i++ {
		res = SampleF2Map(m, n)
	}
	_ = res
}

func BenchmarkSampleP(b *testing.B) {
	var res []int
	for i := 0; i < b.N; i++ {
		res = SampleP(m, n)
	}
	_ = res
}

func BenchmarkSamplePMap(b *testing.B) {
	var res []int
	for i := 0; i < b.N; i++ {
		res = SamplePMap(m, n)
	}
	_ = res
}

func TestSamples(t *testing.T) {
	testRepetitions := 100
	tests := []struct {
		Name     string
		SampleFn func(int, int) []int
	}{
		{
			"Sample",
			Sample,
		},
		{
			"SampleF2",
			SampleF2,
		},
		{
			"SampleF2Map",
			SampleF2Map,
		},
		{
			"SampleP",
			SampleP,
		},
		{
			"SamplePMap",
			SamplePMap,
		},
	}

	for _, test := range tests {
		for i := 0; i < testRepetitions; i++ {
			t.Run(test.Name, func(t *testing.T) {
				t.Parallel()
				res := test.SampleFn(m, n)

				if len(res) != m {
					t.Fatalf("(%s) Unexpected number of elememens in result, want = %d, got = %d", test.Name, m, len(res))
				}

				seen := make(map[int]int, len(res))
				for i, ele := range res {
					if index, ok := seen[ele]; ok {
						t.Fatalf("(%s) Repeated element, first seen at %d, found again at %d", test.Name, index, i)
					}
					seen[ele] = i
				}
			})
		}
	}
}
