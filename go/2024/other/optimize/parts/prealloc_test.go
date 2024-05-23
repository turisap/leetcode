package parts

import "testing"

var length = 50000

//var length = 1000

func BenchmarkPreallocateNo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Don't preallocate our initial slice
		var init [][32]byte
		for j := 0; j < length; j++ {
			init = append(init, [32]byte{})
		}
	}
}

func BenchmarkPreallocate(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Preallocate our initial slice
		init := make([][32]byte, 0, length)
		for j := 0; j < length; j++ {
			init = append(init, [32]byte{})
		}
	}
}
