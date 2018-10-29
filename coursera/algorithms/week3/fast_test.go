package week4

import (
	"io/ioutil"
	"path"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestFastCollinearPoints(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FastCollinearPoints(tt.args.points); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FastCollinearPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}

func testFiles(b *testing.B) []string {
	files, err := ioutil.ReadDir("./testdata")
	if err != nil {
		b.Error(err.Error())
	}
	var names []string

	for _, f := range files {
		if strings.Contains(f.Name(), "input") {
			names = append(names, f.Name())
		}
	}
	return names
}

func loadTest(name string, b *testing.B) []Point {
	f, err := ioutil.ReadFile(path.Join("testdata", name))
	if err != nil {
		b.Error(err.Error())
	}

	lines := strings.Split(string(f), "\n")
	var points []Point
	for _, line := range lines[1:] {
		if line == "" {
			break
		}
		xy := strings.Fields(line)
		x, _ := strconv.Atoi(xy[0])
		y, _ := strconv.Atoi(xy[1])
		points = append(points, Point{x, y})
	}
	return points
}
func BenchmarkFastCollinearPoints_largeInput(b *testing.B) {
	var got []LineSegment
	tests := testFiles(b)
	for _, test := range tests {
		points := loadTest(test, b)
		b.Run(test, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				got = FastCollinearPoints(points)
			}
		})
	}
	_ = got
}
