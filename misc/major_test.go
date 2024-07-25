package misc

import (
	"slices"
	"testing"
)

func TestFindMajorExists(t *testing.T) {
	testsExists := []struct {
		s         []int
		wantMajor int
	}{
		{[]int{1}, 1},
		{[]int{1, 2, 2}, 2},
		{[]int{3, 2, 3}, 3},
		{[]int{2, 2, 1, 1, 1, 2, 2}, 2},
		{[]int{1, 1, 2, 1, 2, 3, 3, 2, 2, 2, 1, 2, 2, 3, 2, 2}, 2},
		{[]int{2, 2, 3, 1, 4, 5, 2, 2, 3, 2, 2, 2, 1, 2, 1, 2, 2, 1, 2, 2}, 2},
	}

	for _, tt := range testsExists {
		major := FindMajor(slices.Values(tt.s))
		if major != tt.wantMajor {
			t.Errorf("FindMajor(%v) = %d, want %d", tt.s, tt.wantMajor, major)
		}
	}
}
