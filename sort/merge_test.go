package sort

import (
	"cmp"
	"slices"
	"testing"
)

var mergeData = []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}

func TestMergeFunc_EmptyNil(t *testing.T) {
	testEmptyNilSlice(t, Merge[[]int])
}

func TestMerge_Data(t *testing.T) {
	testData(t, Merge, mergeData, cmp.Compare)
}

func TestMerge_RandomInts(t *testing.T) {
	testRandomInts(t, Merge, cmp.Compare)
}

func TestMergeFunc_Stability(t *testing.T) {
	n, m := 100000, 1000
	if testing.Short() {
		n, m = 1000, 100
	}
	testStability(t, Merge, n, m)
}

func BenchmarkMerge(b *testing.B) {
	benchmarkInts(b, Merge, 100_000)
}

func FuzzMergeSortFunc(f *testing.F) {
	f.Fuzz(func(t *testing.T, s []byte) {
		Merge(s, cmp.Compare)
		if !slices.IsSortedFunc(s, cmp.Compare) {
			t.Errorf("slice was not sorted")
		}
	})
}
