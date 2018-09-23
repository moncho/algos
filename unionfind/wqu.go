package unionfind

type UnionFind interface {
	Union(a, b int)
	Connected(a, b int) bool
}

type weightedQuickUnion struct {
	roots   []int
	weights []int
}

func NewWeightedQuickUnion(elements int) UnionFind {
	roots := make([]int, elements)
	weights := make([]int, elements)

	for i := range roots {
		roots[i] = i
		weights[i] = 1
	}
	return &weightedQuickUnion{
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

func (qu *weightedQuickUnion) Union(a, b int) {
	rootB := qu.root(b)
	rootA := qu.root(a)
	if rootB == rootA {
		return
	}
	if qu.weights[rootA] > qu.weights[rootB] {
		qu.weights[rootA] += qu.weights[rootB]
		qu.roots[rootB] = rootA
	} else {
		qu.weights[rootB] += qu.weights[rootA]
		qu.roots[rootA] = rootB
	}
}

func (qu *weightedQuickUnion) Connected(a, b int) bool {
	return qu.root(a) == qu.root(b)
}
