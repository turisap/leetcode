package l_word

import "testing"

var res int
var input = `i am traveling down the river
i am traveling down the riveri am traveling down the riveri am traveling down the riveri am traveling down the river
i am traveling down the riveri am traveling down the riveri am traveling down the riveri am traveling down the rive
ri am traveling down the riveri am traveling down the riveri am traveling down the riveri am traveling down the riveri am traveling down the river`

func BenchmarkLastWordNaive(b *testing.B) {
	for _ = range b.N {
		res = lastWordNaive(input)
	}
}

func BenchmarkLastWordIterative(b *testing.B) {
	for _ = range b.N {
		res = lastWordIterative(input)
	}
}
