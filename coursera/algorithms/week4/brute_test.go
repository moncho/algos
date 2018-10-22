package week4

import (
	"reflect"
	"testing"
)

func TestBruteCollinearPoints(t *testing.T) {
	type args struct {
		points []Point
	}
	tests := []struct {
		name string
		args args
		want []LineSegment
	}{
		{
			"",
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BruteCollinearPoints(tt.args.points); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BruteCollinearPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
