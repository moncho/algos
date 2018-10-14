package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/moncho/algos/unionfind"
)

type socialNet struct {
	members    int
	partitions int
	uf         unionfind.UnionFind
}

func newSocialNetwork(members int) *socialNet {
	return &socialNet{
		members:    members,
		partitions: members,
		uf:         unionfind.NewWeightedQuickUnion(members),
	}
}

func (sn *socialNet) NewFriends(a, b int) {
	if sn.uf.Connected(a, b) {
		return
	}
	sn.uf.Union(a, b)
	sn.partitions -= 1
}

func (sn *socialNet) AllFriends() bool {
	return sn.partitions == 1
}

func main() {
	path := os.Args[1]
	log, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	lines := strings.Split(string(log), "\n")

	n, _ := strconv.Atoi(lines[0])
	sn := newSocialNetwork(n)
	var timestamp string
	for _, line := range lines[1:] {
		if line == "" {
			break
		}
		ff := strings.Fields(line)

		friend1, _ := strconv.Atoi(ff[1])
		friend2, _ := strconv.Atoi(ff[2])
		sn.NewFriends(friend1, friend2)
		if sn.AllFriends() {
			timestamp = ff[0]
			break
		}
	}
	if timestamp != "" {
		fmt.Printf("The world became a better place on %v\n", timestamp)
	} else {
		fmt.Println("There is no hope")
	}

	os.Exit(0)
}
