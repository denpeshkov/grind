package unionfind

import (
	"fmt"
	"slices"
	"testing"
)

type uf struct {
	p [][]int
}

func newUf(n int) *uf {
	p := make([][]int, n)
	for i := range p {
		p[i] = []int{i}
	}
	return &uf{p: p}
}

func (u *uf) Union(p, q int) {
	pi, qi := u.ind(p), u.ind(q)
	if pi == qi {
		return
	}
	u.p[pi] = append(u.p[pi], u.p[qi]...)
	u.p = slices.Delete(u.p, qi, qi+1)
}

func (u *uf) Connected(p, q int) bool {
	for _, s := range u.p {
		if slices.Contains(s, p) && slices.Contains(s, q) {
			return true
		}
	}
	return false
}

func (u *uf) ind(p int) int {
	return slices.IndexFunc(u.p, func(s []int) bool {
		return slices.Contains(s, p)
	})
}

func TestUnionFind(t *testing.T) {
	tests := []struct {
		n int
		u [][2]int
	}{
		{
			n: 1,
			u: [][2]int{{0, 0}, {0, 0}},
		},
		{
			n: 2,
			u: [][2]int{{0, 0}, {0, 0}, {0, 1}, {0, 1}, {1, 1}, {1, 0}, {1, 0}, {1, 1}},
		},
		{
			n: 3,
			u: [][2]int{{0, 0}, {0, 0}, {0, 1}, {0, 1}, {1, 2}, {0, 2}},
		},
		{
			n: 5,
			u: [][2]int{{0, 1}, {1, 1}, {2, 2}, {2, 3}, {2, 4}, {2, 4}, {0, 4}, {3, 4}, {3, 4}},
		},
		{
			n: 10,
			u: [][2]int{{0, 0}, {0, 9}, {5, 4}, {4, 5}, {5, 5}, {7, 7}, {7, 0}, {7, 5}},
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("n=%d", tt.n), func(t *testing.T) {
			t.Parallel()

			uft := newUf(tt.n)
			uf := New(tt.n)
			for _, u := range tt.u {
				p, q := u[0], u[1]
				t.Logf("UnionFind.Union(%v, %v)", p, q)
				uf.Union(p, q)
				uft.Union(p, q)
				testUFConnected(t, uf, uft)

				pb := make([]bool, tt.n)
				for p, sz := range uf.All() {
					if uft.ind(p) == -1 {
						t.Errorf("UnionFind.All() returned wrong representative %d", p)
					}
					if l := len(uft.p[uft.ind(p)]); sz != l {
						t.Errorf("UnionFind.All() returned size %d, want %d", sz, l)
					}
					pb[p] = true
				}
				cnt := 0
				for _, b := range pb {
					if b {
						cnt++
					}
				}
				if cnt != len(uft.p) {
					t.Errorf("UnionFind.All() returned %d representatives, want %d", cnt, len(uft.p))
				}

				if uf.Count() != cnt {
					t.Errorf("UnionFind.Count() = %d, want %d", uf.Count(), len(uft.p))
				}
			}
		})
	}
}

func testUFConnected(t *testing.T, uf *UnionFind, uft *uf) {
	for p := range len(uf.parent) {
		for q := range len(uf.parent) {
			got, want := uf.Connected(p, q), uft.Connected(p, q)
			if got != want {
				t.Errorf("UnionFind.Connected(%d, %d) = %t, want %t", p, q, got, want)
			}
		}
	}
}
