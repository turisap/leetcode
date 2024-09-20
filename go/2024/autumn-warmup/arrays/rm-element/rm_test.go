package rm_element

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tcc = []struct {
	nums         []int
	expectedNums []int
	val          int
	out          int
}{
	{
		nums:         []int{3, 2, 2, 3},
		val:          3,
		expectedNums: []int{2, 2, 0, 0},
		out:          2,
	},
	{
		nums:         []int{0, 1, 2, 2, 3, 0, 4, 2},
		val:          2,
		expectedNums: []int{0, 1, 3, 0, 4, 0, 0, 0},
		out:          5,
	},
}

func TestNaiveRemove(t *testing.T) {
	for _, tc := range tcc {
		res := removeElementNaive(tc.nums, tc.val)
		assert.Equal(t, tc.out, res)
		assert.ElementsMatch(t, tc.expectedNums, tc.nums)
	}
}

//func TestPointersRemove(t *testing.T) {
//	for _, tc := range tcc {
//		res := removeElementsPointers(tc.nums, tc.val)
//		assert.Equal(t, tc.out, res)
//		assert.ElementsMatch(t, tc.expectedNums, tc.nums)
//	}
//}
