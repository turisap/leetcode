package lcp

import "testing"

var r string
var i = []string{
	"flapple", "flamazing", "flamingo", "flapjack", "flavor", "flash", "flask", "flair", "flannel", "flange",
	"flap", "flapper", "flattop", "flatiron", "flavorful", "flavoring", "flawed", "flaunt", "flax", "flame",
	"flabbergasted", "flabby", "flameproof", "flaming", "flank", "flame-retardant", "flamethrower", "flak",
	"flatulent", "flatware", "flawless", "flaky", "flaxen", "flavorsome", "flamboyant", "flautist", "flaccid",
	"flamingos", "flask-full", "flawed-edge", "flak-proof", "flaming-red", "flank-steak", "flavor-blast",
	"flashbang", "flashlight", "flash-freeze", "flab-attack", "flat-screen", "flaming-cheetah",
}

func BenchmarkLCPNaive(b *testing.B) {
	for _ = range b.N {
		r = longestCommonPrefixNaive(i)
	}
}
func BenchmarkLCPPresort(b *testing.B) {
	for _ = range b.N {
		r = longestCommonPresort(i)
	}
}
