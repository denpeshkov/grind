package sort

// Insertion sorts the slice s in ascending order as determined by the cmp function using top-down merge sort.
//
// Time: O(n*log(n))
// Space: O(n)
func Merge[S ~[]E, E any](s S, cmp func(a, b E) int) {
	aux := make(S, len(s))
	var sort func(s S)
	sort = func(s S) {
		if len(s) <= 1 {
			return
		}
		m := len(s) / 2
		sort(s[:m])
		sort(s[m:])
		merge(s[:m], s[m:], cmp, aux)
		copy(s, aux)
	}
	sort(s)
}

func merge[S ~[]E, E any](a, b S, cmp func(a, b E) int, aux S) {
	for i, j, k := 0, 0, 0; k < len(a)+len(b); k++ {
		// ensures stability
		if i >= len(a) || (j < len(b) && cmp(b[j], a[i]) < 0) {
			aux[k] = b[j]
			j++
		} else {
			aux[k] = a[i]
			i++
		}
	}
}
