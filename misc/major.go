package misc

import (
	"iter"
)

// FindMajor finds the majority element in a sequence of elements.
// It returns the majority element if one exists; otherwise, it returns an arbitrary element.
// It is a streaming algorithm requiring only a single pass through the iterator.
// It uses Boyer–Moore majority vote streaming algorithm.
//
// Time: O(N)
// Space: O(1)
func FindMajor[T comparable](seq iter.Seq[T]) T {
	var (
		m T
		c int
	)
	for v := range seq {
		if c == 0 {
			m = v
		}
		if m == v {
			c++
		} else {
			c--
		}
	}
	return m
}
