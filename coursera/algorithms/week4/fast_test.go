package week4

import (
	"reflect"
	"testing"
)

func TestFastCollinearPoints(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FastCollinearPoints(tt.args.points); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BruteCollinearPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
