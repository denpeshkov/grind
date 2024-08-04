package sort

import (
	"cmp"
	"math/rand"
	"slices"
	"sort"
	"testing"
)

func testEmptyNilSlice[S ~[]E, E any](t *testing.T, f func(S, func(a, b E) int)) {
	emptySlice := []E{}
	var nilSlice []E

	panics := func(f func()) (b bool) {
		defer func() {
			if x := recover(); x != nil {
				b = true
			}
		}()
		f()
		return false
	}

	if panics(func() { f(emptySlice, func(a, b E) int { return 0 }) }) {
		t.Errorf("got unexpected panic on empty slice")
	}
	if panics(func() { f(nilSlice, func(a, b E) int { return 0 }) }) {
		t.Errorf("got unexpected panic on nil slice")
	}
}

func testData[S ~[]E, E any](t *testing.T, f func(S, func(a, b E) int), s S, cmp func(a, b E) int) {
	s = slices.Clone(s)
	f(s, cmp)
	if !slices.IsSortedFunc(s, cmp) {
		t.Errorf("got unsorted slice: %v, want sorted slice", s)
	}
}

func testRandomInts(t *testing.T, f func([]int, func(a, b int) int), cmp func(a, b int) int) {
	n := 1000
	s := make([]int, n)
	for i := 0; i < len(s); i++ {
		s[i] = rand.Intn(1000)
	}
	if slices.IsSorted(s) {
		t.Fatal("terrible rand")
	}

	f(s, cmp)
	if !slices.IsSortedFunc(s, cmp) {
		t.Errorf("sort didn't sort - %d random ints", n)
	}
}

type intPair struct {
	a, b int
}

// Pairs compare on a only.
func intPairCmp(x, y intPair) int {
	return x.a - y.a
}

type intPairs []intPair

func (d intPairs) Len() int           { return len(d) }
func (d intPairs) Less(i, j int) bool { return intPairCmp(d[i], d[j]) < 0 }
func (d intPairs) Swap(i, j int)      { d[i], d[j] = d[j], d[i] }

// inOrder checks if a-equal elements were not reordered.
func (d intPairs) inOrder() bool {
	lastA, lastB := -1, 0
	for i := 0; i < len(d); i++ {
		if lastA != d[i].a {
			lastA = d[i].a
			lastB = d[i].b
			continue
		}
		if d[i].b <= lastB {
			return false
		}
		lastB = d[i].b
	}
	return true
}

// initB records initial order in B.
func (d intPairs) initB() {
	for i := range d {
		d[i].b = i
	}
}

func testStability(t *testing.T, f func(intPairs, func(a, b intPair) int), n, m int) {
	data := make(intPairs, n)

	// random distribution
	for i := 0; i < len(data); i++ {
		data[i].a = rand.Intn(m)
	}
	if sort.IsSorted(data) {
		t.Fatalf("terrible rand.rand")
	}
	data.initB()
	f(data, intPairCmp)
	if !data.inOrder() {
		t.Errorf("sort wasn't stable on %d ints", n)
	}

	// already sorted
	data.initB()
	f(data, intPairCmp)
	if !data.inOrder() {
		t.Errorf("sort shuffled sorted %d ints (stability)", n)
	}

	// sorted reversed
	for i := 0; i < len(data); i++ {
		data[i].a = len(data) - i
	}
	data.initB()
	f(data, intPairCmp)
	if !data.inOrder() {
		t.Errorf("sort wasn't stable on %d ints", n)
	}
}

func benchmarkInts(b *testing.B, sortFunc func([]int, func(a, b int) int), sz int) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		data := make([]int, sz)
		for i := 0; i < len(data); i++ {
			data[i] = i ^ 0x2cc
		}
		b.StartTimer()
		sortFunc(data, cmp.Compare)
		b.StopTimer()
	}
}
