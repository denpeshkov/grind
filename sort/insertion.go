package sort

import (
	"sort"
)

// Insertion sorts the slice s in ascending order as determined by the cmp function using insertion sort.
//
// Time: O(n^2)
// Space: O(1)
func Insertion[S ~[]E, E any](s S, cmp func(a, b E) int) {
	InsertionD(&slice[E]{s: s, cmp: cmp})
}

// InsertionD sorts data in ascending order as determined by [sort.Interface] using insertion sort.
//
// Time: O(n^2)
// Space: O(1)
func InsertionD(data sort.Interface) {
	for i := 1; i < data.Len(); i++ {
		for j := i; j > 0 && data.Less(j, j-1); j-- {
			data.Swap(j, j-1)
		}
	}
}
