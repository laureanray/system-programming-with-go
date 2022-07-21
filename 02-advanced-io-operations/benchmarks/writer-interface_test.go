package benchmarks

import (
	"advanced-io-operations/impl"
	"testing"
)

func BenchmarkReverseWrite(b *testing.B) {
	for i := 0; i < b.N; i++ {
		impl.ReverseWrite("../data/test.txt", "../data/dest.txt")
	}
}
