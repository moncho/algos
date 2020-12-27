package main

import (
	"math/rand"
)

// Sample takes m samples from 1 to n.
func Sample(m, n int) []int {
	if m == 0 {
		return nil
	}

	s := Sample(m-1, n-1)
	t := rand.Intn(n) + 1
	if contains(t, s) {
		s = append(s, n)
	} else {
		s = append(s, t)
	}
	return s
}

// SampleF2 takes m samples from 1 to n.
func SampleF2(m, n int) []int {
	res := make([]int, m)
	i := 0
	for j := n - m + 1; j <= n; j++ {
		t := rand.Intn(j) + 1
		if contains(t, res) {
			res[i] = j
		} else {
			res[i] = t
		}
		i++
	}
	return res
}

// SampleF2Map SampleF2 using maps.
func SampleF2Map(m, n int) []int {
	sample := make(map[int]struct{}, m)
	res := make([]int, m)
	i := 0

	for j := n - m + 1; j <= n; j++ {
		t := rand.Intn(j) + 1
		if _, ok := sample[t]; ok {
			sample[j] = struct{}{}
			res[i] = j
		} else {
			sample[t] = struct{}{}
			res[i] = t
		}
		i++
	}

	return res
}

// SampleP takes m samples from 1 to n in random order.
func SampleP(m, n int) []int {
	res := make([]int, m)
	for j := n - m + 1; j <= n; j++ {
		t := rand.Intn(j) + 1
		if i := indexOf(t, res); i != -1 {
			copy(res[i+1:], res[i:])
			res[i] = j
		} else {
			if len(res) > 0 {
				copy(res[1:], res)
			}
			res[0] = t
		}
	}
	return res
}

// SamplePMap takes m samples from 1 to n in random order.
// This uses Go random key order when ranging o a map to achieve
// the desired random order in samples.
func SamplePMap(m, n int) []int {
	sample := make(map[int]struct{}, m)

	for j := n - m + 1; j <= n; j++ {
		t := rand.Intn(j) + 1
		if _, ok := sample[t]; ok {
			sample[j] = struct{}{}
		} else {
			sample[t] = struct{}{}
		}
	}

	res := make([]int, m, m)
	i := 0
	for k := range sample {
		res[i] = k
		i++
	}

	return res
}
func contains(ele int, ii []int) bool {
	for _, i := range ii {
		if ele == i {
			return true
		}
	}
	return false
}

func indexOf(ele int, ii []int) int {
	for index, i := range ii {
		if ele == i {
			return index
		}
	}
	return -1
}
