package main

import (
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
)

type coord struct {
	row, col int
}
type percolationTest struct {
	n         int
	openSites []coord
}

func TestPercolation(t *testing.T) {
	for _, tc := range testCases(t.Error) {
		percolates := !strings.Contains(tc, "no")
		test := loadTestCase(t.Error, tc)
		p := NewPercolation(test.n)
		for _, open := range test.openSites {
			p.open(open.row, open.col)
		}
		got := p.percolates()
		if percolates != got {
			t.Errorf("Test for %s failed, percolates %t", tc, got)
		}
	}
}
func testCases(f func(...interface{})) []string {
	files, err := ioutil.ReadDir("testdata")
	if err != nil {
		f(err)
	}
	var testCases []string
	for _, file := range files {
		testCases = append(testCases, file.Name())
	}
	return testCases
}

func loadTestCase(onFail func(args ...interface{}), fileName string) percolationTest {
	path := filepath.Join("testdata", fileName)
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		onFail(err)
	}
	lines := strings.Split(string(bytes), "\n")
	n, _ := strconv.Atoi(lines[0])
	var openSites []coord
	for _, line := range lines[1:] {
		if line == "" {
			break
		}
		rowCol := strings.Fields(line)
		row, _ := strconv.Atoi(rowCol[0])
		col, _ := strconv.Atoi(rowCol[1])
		openSites = append(openSites, coord{row, col})
	}
	return percolationTest{n, openSites}
}

var percolates bool

func BenchmarkPercolation(b *testing.B) {
	for _, tc := range testCases(b.Error) {
		tc := tc
		test := loadTestCase(b.Error, tc)
		b.ResetTimer()
		b.Run(tc, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				p := NewPercolation(test.n)
				for _, open := range test.openSites {
					p.open(open.row, open.col)
				}
				percolates = p.percolates()
			}
		})
	}
}
