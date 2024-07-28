package unionfind

import (
	"iter"
)

// UnionFind represents a union–find (disjoint-set) data structure.
// Access to elements out of range causes panic.
//
// Implementation is a Weighted Quick-Union with Path Halving.
//
// Space: O(n)
type UnionFind struct {
	parent []int
	size   []int
	cnt    int
}

// New constructs new UnionFind of with n elements 0 through n-1, each in its own set.
//
// Time: O(n)
func New(n int) *UnionFind {
	p := make([]int, n)
	sz := make([]int, n)
	for i := range n {
		p[i], sz[i] = i, 1
	}
	return &UnionFind{parent: p, size: sz, cnt: n}
}

// Find returns the representative of the set containing element p.
//
// Time: O(α(n)) amortized; where α(n) is an inverse Ackermann function.
func (u *UnionFind) Find(p int) int {
	for p != u.parent[p] {
		u.parent[p] = u.parent[u.parent[p]] // path compression by halving
		p = u.parent[p]
	}
	return p
}

// Connected determines whether two elements are in the same set.
//
// Time: O(α(n)) amortized; where α(n) is an inverse Ackermann function.
func (u *UnionFind) Connected(p, q int) bool {
	return u.Find(p) == u.Find(q)
}

// Union merges the set containing element p with the set containing element q.
//
// Time: O(α(n)) amortized; where α(n) is an inverse Ackermann function.
func (u *UnionFind) Union(p, q int) {
	p, q = u.Find(p), u.Find(q)
	if p == q {
		return
	}
	switch psz, qsz := u.size[p], u.size[q]; {
	case psz <= qsz:
		u.parent[p] = q
		u.size[q] += psz
	case psz > qsz:
		u.parent[q] = p
		u.size[p] += qsz
	}
	u.cnt--
}

// Count returns the number of sets.
//
// Time: O(1)
func (u *UnionFind) Count() int {
	return u.cnt
}

// All returns an iterator over representatives of sets.
//
// Time: O(n)
func (u *UnionFind) All() iter.Seq2[int, int] {
	return func(yield func(int, int) bool) {
		for i, p := range u.parent {
			if p == i && !yield(p, u.size[i]) {
				return
			}
		}
	}
}
