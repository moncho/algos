package main

import (
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
)

func TestPercolation(t *testing.T) {
	for _, tc := range testCases(t) {
		percolates := !strings.Contains(tc, "no")
		test := testCase(t, tc)
		lines := strings.Split(test, "\n")
		n, _ := strconv.Atoi(lines[0])
		p := NewPercolation(n)
		for _, line := range lines[1:] {
			if line == "" {
				break
			}
			rowCol := strings.Fields(line)
			row, _ := strconv.Atoi(rowCol[0])
			col, _ := strconv.Atoi(rowCol[1])

			p.open(row, col)
		}
		got := p.percolates()
		if percolates != got {
			t.Errorf("Test for %s failed, percolates %t", tc, got)
		}
	}
}
func testCases(t *testing.T) []string {
	files, err := ioutil.ReadDir("testdata")
	if err != nil {
		t.Fatal(err)
	}
	var testCases []string
	for _, file := range files {
		testCases = append(testCases, file.Name())
	}
	return testCases
}

func testCase(t *testing.T, name string) string {
	path := filepath.Join("testdata", name)
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	return string(bytes)
}
