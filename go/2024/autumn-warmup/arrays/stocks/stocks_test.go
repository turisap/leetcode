package stocks

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStockCalc(t *testing.T) {
	tcc := []struct {
		input []int
		res   int
	}{
		{
			input: []int{7, 1, 5, 3, 6, 4},
			res:   5,
		},
		{
			input: []int{7, 6, 4, 3, 1},
			res:   0,
		},
	}

	for _, tc := range tcc {
		res := stocksProfit(tc.input)
		assert.Equal(t, tc.res, res)
	}
}
