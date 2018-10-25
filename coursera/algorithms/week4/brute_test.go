package week4

import (
	"reflect"
	"testing"
)

type args struct {
	points []Point
}

var tests = []struct {
	name string
	args args
	want []LineSegment
}{
	{
		"input6.txt",
		args{
			[]Point{
				Point{19000, 10000},
				Point{18000, 10000},
				Point{32000, 10000},
				Point{21000, 10000},
				Point{1234, 5678},
				Point{14000, 10000},
			},
		},
		[]LineSegment{
			LineSegment{
				Point{14000, 10000},
				Point{32000, 10000},
			},
		},
	},
	{
		"input8.txt",
		args{
			[]Point{
				Point{10000, 0},
				Point{0, 10000},
				Point{3000, 7000},
				Point{7000, 3000},
				Point{20000, 21000},
				Point{3000, 4000},
				Point{14000, 15000},
				Point{6000, 7000},
			},
		},
		[]LineSegment{
			LineSegment{
				Point{10000, 0},
				Point{0, 10000},
			},
			LineSegment{
				Point{3000, 4000},
				Point{20000, 21000},
			},
		},
	},
}

func TestBruteCollinearPoints(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BruteCollinearPoints(tt.args.points); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BruteCollinearPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
