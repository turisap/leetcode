package rm_duplicates

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveDuplicatesNaive(t *testing.T) {
	tcc := []struct {
		nums         []int
		uniqueCount  int
		expectedNums []int
	}{
		{
			nums:         []int{1, 2, 2},
			expectedNums: []int{1, 2, 0},
			uniqueCount:  2,
		},
		{
			nums:         []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			expectedNums: []int{0, 1, 2, 3, 4, 0, 0, 0, 0, 0},
			uniqueCount:  5,
		},
	}

	for _, tc := range tcc {
		res := removeDuplicatesNaive(tc.nums)
		assert.Equal(t, tc.uniqueCount, res)
		assert.ElementsMatch(t, tc.nums, tc.expectedNums)
	}
}
