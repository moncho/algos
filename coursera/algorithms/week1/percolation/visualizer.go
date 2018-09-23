package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
)

var (
	open  = color.New(color.FgWhite).SprintFunc()
	full  = color.New(color.FgCyan).SprintFunc()
	close = color.New(color.FgHiBlack).SprintFunc()
)

const block = string('\u2589')

func main() {
	if len(os.Args) < 2 {
		fmt.Println("a file path is expected")
		os.Exit(1)
	}
	bytes, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("error reading %s: %v\n", os.Args[1], err)
		os.Exit(1)
	}
	content := string(bytes)
	lines := strings.Split(content, "\n")
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
		draw(p, n)
		time.Sleep(200 * time.Millisecond)
	}

}

func draw(perc Percolation, n int) {
	//Reset the terminal
	fmt.Printf("\033[2J\033[H")
	var buffer strings.Builder
	for row := 1; row <= n; row++ {
		for col := 1; col <= n; col++ {
			printBlock := close
			if perc.isFull(row, col) {
				printBlock = full
			} else if perc.isOpen(row, col) {
				printBlock = open
			}
			fmt.Fprintf(&buffer, "%s ", printBlock(block))
		}
		fmt.Fprint(&buffer, "\n")
	}
	percolates := perc.percolates()
	fmt.Fprintf(&buffer, "%d open sites\n", perc.numberOfOpenSites())
	if percolates {
		fmt.Fprintln(&buffer, "It percolates!!")
	} else {
		fmt.Fprintln(&buffer, "It does not percolate")
	}
	fmt.Printf("%s", buffer.String())
}
