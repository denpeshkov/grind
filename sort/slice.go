package sort

// slice implements [sort.Interface] on a generic slice.
type slice[E any] struct {
	s   []E
	cmp func(a, b E) int
}

func (s *slice[E]) Len() int               { return len(s.s) }
func (s *slice[E]) Less(i int, j int) bool { return s.cmp(s.s[i], s.s[j]) < 0 }
func (s *slice[E]) Swap(i int, j int)      { s.s[i], s.s[j] = s.s[j], s.s[i] }
