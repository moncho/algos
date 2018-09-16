package unionfind

type UnionFind interface {
	union(a, b int)
	connected(a, b int) bool
}

type weightedQuickUnion struct {
	roots   []int
	weights []int
}

func NewWeightedQuickUnion(elements int) weightedQuickUnion {
	roots := make([]int, elements)
	weights := make([]int, elements)

	for i := range roots {
		roots[i] = i
		weights[i] = 1
	}
	return weightedQuickUnion{
		roots:   roots,
		weights: weights,
	}
}

func (qu *weightedQuickUnion) root(a int) int {
	root := qu.roots[a]
	for root != qu.roots[root] {
		//flatten tree
		qu.roots[root] = qu.roots[qu.roots[root]]
		root = qu.roots[root]
	}
	return root
}

func (qu *weightedQuickUnion) union(a, b int) {
	rootB := qu.root(b)
	rootA := qu.root(b)
	if qu.weights[rootA] > qu.weights[rootB] {
		qu.weights[rootA] += qu.weights[rootB]
		qu.roots[b] = rootA
	} else {
		qu.weights[rootB] += qu.weights[rootA]
		qu.roots[a] = rootB
	}
}

func (qu *weightedQuickUnion) connected(a, b int) bool {
	return qu.root(a) == qu.root(b)
}
