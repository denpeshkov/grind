package unionfind

import (
	"fmt"
	"testing"
)

func TestElement(t *testing.T) {
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
			uft := newUf(tt.n)
			elems := make([]*Element[int], tt.n)
			for i := range tt.n {
				e := &Element[int]{Value: i}
				MakeSet(e)
				elems[i] = e
			}

			for _, e := range elems {
				if got := Find(e); got != e {
					t.Fatalf("Find(%v) = %v, want %v", e, got, e)
				}
			}

			for _, u := range tt.u {
				p, q := u[0], u[1]
				Union(elems[p], elems[q])
				uft.Union(p, q)
				testConnected(t, elems, uft)
			}
		})
	}
}

func TestFindMakeSet(t *testing.T) {
	e := &Element[int]{Value: 1}
	if got := Find(e); got != nil {
		t.Fatalf("Find() before MakeSet() returned non nil")
	}

	MakeSet(e)
	if got := Find(e); got != e {
		t.Fatalf("Find(%v) = %v, want %v", e, got, e)
	}
}

func TestMakeSetFind(t *testing.T) {
	e1 := &Element[int]{Value: 1}
	MakeSet(e1)
	e2 := &Element[int]{Value: 2}
	MakeSet(e2)

	Union(e1, e2)
	if !Connected(e1, e2) {
		t.Fatal("Union() didn't connect")
	}

	MakeSet(e1)
	if !Connected(e1, e2) {
		t.Errorf("MakeSet(%v) disconnected existing element", e1)
	}
	MakeSet(e2)
	if !Connected(e1, e2) {
		t.Errorf("MakeSet(%v) disconnected existing element", e2)
	}
}

func testConnected(t *testing.T, els []*Element[int], uft *uf) {
	for p := range len(els) {
		for q := range len(els) {
			got, want := Connected(els[p], els[q]), uft.Connected(p, q)
			if got != want {
				t.Errorf("Connected(%v, %v) = %t, want %t", p, q, got, want)
			}
		}
	}
}
