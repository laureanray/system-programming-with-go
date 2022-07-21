package benchmarks

import (
	"advanced-io-operations/impl"
	"testing"
)

func BenchmarkPeeking(b *testing.B) {
	for i := 0; i < b.N; i++ {
		impl.Peek("../data/pypi.index")
	}
}
