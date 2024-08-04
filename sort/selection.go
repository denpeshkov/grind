package sort

import (
	"sort"
)

// Selection sorts the slice s in ascending order as determined by the cmp function using selection sort.
//
// Time: O(n^2)
// Space: O(1)
func Selection[S ~[]E, E any](s S, cmp func(a, b E) int) {
	SelectionD(&slice[E]{s: s, cmp: cmp})
}

// SelectionD sorts data in ascending order as determined by [sort.Interface] using selection sort.
//
// Time: O(n^2)
// Space: O(1)
func SelectionD(data sort.Interface) {
	for i := range data.Len() {
		mini := i
		for j := i + 1; j < data.Len(); j++ {
			if data.Less(j, mini) {
				mini = j
			}
		}
		data.Swap(i, mini)
	}
}
