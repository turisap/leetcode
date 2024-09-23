package lcp

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tcc = []struct {
	input []string
	out   string
}{
	{
		input: []string{"flower", "flow", "flight"},
		out:   "fl",
	},
	{
		input: []string{"dog", "racecar", "car"},
		out:   "",
	},
	{
		input: []string{"ab", "a"},
		out:   "a",
	},
	{
		input: []string{"reflower", "flow", "flight"},
		out:   "",
	},
	{
		input: []string{"aaa", "aa", "aaa"},
		out:   "aa",
	},
	{
		input: []string{
			"flapple", "flamazing", "flamingo", "flapjack", "flavor", "flash", "flask", "flair", "flannel", "flange",
			"flap", "flapper", "flattop", "flatiron", "flavorful", "flavoring", "flawed", "flaunt", "flax", "flame",
			"flabbergasted", "flabby", "flameproof", "flaming", "flank", "flame-retardant", "flamethrower", "flak",
			"flatulent", "flatware", "flawless", "flaky", "flaxen", "flavorsome", "flamboyant", "flautist", "flaccid",
			"flamingos", "flask-full", "flawed-edge", "flak-proof", "flaming-red", "flank-steak", "flavor-blast",
			"flashbang", "flashlight", "flash-freeze", "flab-attack", "flat-screen", "flaming-cheetah",
		},
		out: "fla",
	},
}

func TestLCPNaive(t *testing.T) {
	for _, tc := range tcc {
		res := longestCommonPrefixNaive(tc.input)
		assert.Equal(t, tc.out, res)
	}
}

func TestLCPRegular(t *testing.T) {
	for _, tc := range tcc {
		res := longestCommonPresort(tc.input)
		assert.Equal(t, tc.out, res)
	}
}
