package benchmarks

import (
	"advanced-io-operations/impl"
	"testing"
)

func BenchmarkFileStructure(b *testing.B) {
	for i := 0; i < b.N; i++ {
		impl.Load("../data/pypi.index")
	}
}
