package ransomnote

import "testing"

var res bool
var input = "aabcdgefteriajlkdnkdfghhkncbljklshouqwkrjknbknxjfkfjsmlkdfjlkdjhgksdsdlfjsfkjlsdfjxksjdlfjdslkfjz"
var needle = "axz"

func BenchmarkConstructNaive(b *testing.B) {
	for _ = range b.N {
		res = canConstructNaive(input, needle)
	}
}

func BenchmarkConstructNaiveImproved(b *testing.B) {
	for _ = range b.N {
		res = canConstructNaiveImproved(input, needle)
	}
}

func BenchmarkConstructOneCount(b *testing.B) {
	for _ = range b.N {
		res = canConstructOneCount(input, needle)
	}
}
