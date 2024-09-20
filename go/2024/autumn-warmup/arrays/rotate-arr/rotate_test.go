package rotate_arr

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRotateNaive(t *testing.T) {
	tcc := []struct {
		nums     []int
		k        int
		expected []int
	}{
		{
			nums:     []int{1, 2, 3, 4, 5, 6, 7},
			k:        3,
			expected: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			nums:     []int{-1, -100, 3, 99},
			k:        2,
			expected: []int{3, 99, -1, -100},
		},
	}

	for _, tc := range tcc {
		r := rotateNaive(tc.nums, tc.k)
		assert.Equal(t, tc.expected, r)
	}
}
