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
		expectedNums: []int{2, 2},
		out:          2,
	},
}

func TestNaiveRemove(t *testing.T) {
	for _, tc := range tcc {
		res := removeElementNaive(tc.nums, tc.val)
		assert.Equal(t, tc.out, res)
	}
}
