package unionfind

import "fmt"

// Element represents an element in the set with a given value.
// Each instance of Element is a distinct member of the set, even if they have the same value.
type Element[T any] struct {
	Value  T
	parent *Element[T]
	sz     int
}

func (e *Element[T]) String() string {
	return fmt.Sprintf("%v", e.Value)
}

// MakeSet creates a new set whose only member (and thus representative) is e.
// If e is already an element of one of the sets, e is not modified.
//
// Time: O(1)
func MakeSet[T any](e *Element[T]) {
	if e.parent != nil {
		return
	}
	e.parent, e.sz = e, 1
}

// Find returns the representative of the set containing element e.
// If e is not element of any set, it returns nil.
//
// Time: O(α(n)) amortized; where α(n) is an inverse Ackermann function.
func Find[T any](e *Element[T]) *Element[T] {
	if e.parent == nil {
		return nil
	}
	for e.parent != e {
		e.parent = e.parent.parent // path compression by halving
		e = e.parent
	}
	return e
}

// Connected determines whether two elements are in the same set.
//
// Time: O(α(n)) amortized; where α(n) is an inverse Ackermann function.
func Connected[T any](p, q *Element[T]) bool {
	return Find(p) == Find(q)
}

// Union merges the set containing element p with the set containing element q.
// If either p or q is not part of any set, no modifications are made.
//
// Time: O(α(n)) amortized; where where α(n) is an inverse Ackermann function.
func Union[T any](p, q *Element[T]) {
	p, q = Find(p), Find(q)
	if p == q {
		return
	}
	switch {
	case p.sz <= q.sz:
		p.parent = q
		q.sz += p.sz
	case p.sz > q.sz:
		q.parent = p
		p.sz += q.sz
	}
}
