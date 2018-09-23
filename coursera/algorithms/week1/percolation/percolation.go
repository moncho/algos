package main

import (
	"github.com/moncho/algos/unionfind"
)

type Percolation interface {
	open(row, col int)            // open site (row, col) if it is not open already
	isOpen(row int, col int) bool // is site (row, col) open?
	isFull(row int, col int) bool // is site (row, col) full?
	numberOfOpenSites() int       // number of open sites
	percolates() bool             // does the system percolate?
}

func NewPercolation(n int) Percolation {
	return &perc{
		qu:            unionfind.NewWeightedQuickUnion(n * n),
		n:             n,
		openSites:     make(map[int]bool),
		rowColToIndex: rowColToIndexFunc(n),
	}
}

type perc struct {
	qu            unionfind.UnionFind
	openSites     map[int]bool
	n             int
	rowColToIndex func(int, int) int
}

func (p *perc) open(row, col int) {
	i := p.rowColToIndex(row, col)
	p.openSites[i] = true
	for _, n := range p.neighbours(row, col) {
		p.qu.Union(i, n)
	}
}
func (p *perc) isOpen(row int, col int) bool {
	return p.openSites[p.rowColToIndex(row, col)]
}
func (p *perc) isFull(row int, col int) bool {
	i := p.rowColToIndex(row, col)
	if !p.openSites[i] {
		return false
	}
	if row == 1 {
		return true
	}
	for firstRowCol := 1; firstRowCol <= p.n; firstRowCol++ {
		if p.qu.Connected(p.rowColToIndex(1, firstRowCol), i) {
			return true
		}
	}
	return false
}
func (p *perc) numberOfOpenSites() int {
	return len(p.openSites)
}
func (p *perc) percolates() bool {
	lastrow := p.n
	for col := 1; col <= p.n; col++ {
		if p.isFull(lastrow, col) {
			return true
		}
	}
	return false
}

func (p *perc) neighbours(row, col int) []int {
	var result []int

	if row-1 >= 1 {
		i := p.rowColToIndex(row-1, col)
		if p.openSites[i] {
			result = append(result, i)
		}
	}
	if col-1 >= 1 {
		i := p.rowColToIndex(row, col-1)
		if p.openSites[i] {
			result = append(result, i)
		}
	}
	if row+1 <= p.n {
		i := p.rowColToIndex(row+1, col)
		if p.openSites[i] {
			result = append(result, i)
		}
	}
	if col+1 <= p.n {
		i := p.rowColToIndex(row, col+1)
		if p.openSites[i] {
			result = append(result, i)
		}
	}

	return result
}

func rowColToIndexFunc(width int) func(int, int) int {
	return func(row, col int) int {
		return width*(row-1) + (col - 1)
	}
}
