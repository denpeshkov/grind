package sort

import (
	"cmp"
	"slices"
	"testing"
)

var insertionData = []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}

func TestInsertion_EmptyNil(t *testing.T) {
	testEmptyNilSlice(t, Insertion[[]int])
}

func TestInsertion_Data(t *testing.T) {
	testData(t, Insertion, insertionData, cmp.Compare)
}

func TestInsertion_RandomInts(t *testing.T) {
	testRandomInts(t, Insertion, cmp.Compare)
}

func TestInsertion_Stability(t *testing.T) {
	testStability(t, Insertion, 1000, 100)
}

func BenchmarkInsertion(b *testing.B) {
	benchmarkInts(b, Insertion, 10_000)
}

func FuzzInsertion(f *testing.F) {
	f.Fuzz(func(t *testing.T, s []byte) {
		Insertion(s, cmp.Compare)
		if !slices.IsSorted(s) {
			t.Errorf("slice was not sorted")
		}
	})
}
