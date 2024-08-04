package sort

import (
	"cmp"
	"slices"
	"testing"
)

var selectionData = []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}

func TestSelection_EmptyNil(t *testing.T) {
	testEmptyNilSlice(t, Selection[[]int])
}

func TestSelection_Data(t *testing.T) {
	testData(t, Selection, selectionData, cmp.Compare)
}

func TestSelection_RandomInts(t *testing.T) {
	testRandomInts(t, Selection, cmp.Compare)
}

func BenchmarkSelection(b *testing.B) {
	benchmarkInts(b, Selection, 10_000)
}

func FuzzSelectionFunc(f *testing.F) {
	f.Fuzz(func(t *testing.T, s []byte) {
		Selection(s, cmp.Compare)
		if !slices.IsSorted(s) {
			t.Errorf("slice was not sorted")
		}
	})
}
