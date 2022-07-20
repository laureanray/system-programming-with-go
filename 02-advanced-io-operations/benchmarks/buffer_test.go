package benchmarks

import (
	"advanced-io-operations/impl"
	"testing"
)

func BenchmarkUsingBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		impl.UseBuffer()
	}
}
