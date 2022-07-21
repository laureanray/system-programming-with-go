package benchmarks

import (
	"advanced-io-operations/impl"
	"testing"
)

func BenchmarkFormatWrite(b *testing.B) {
	for i := 0; i < b.N; i++ {
		impl.FormatWrite("../data/format_write.txt")
	}
}
