package benchmarks

import (
	"advanced-io-operations/impl"
	"io/ioutil"
	"strings"
	"testing"
)

func BenchmarkCopyN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r := strings.NewReader("This is an example of CopyN wiht offset")
		for j, l, step := int64(0), int64(r.Len()), int64(5); j < l; j += step {
			impl.CopyNOffset(ioutil.Discard, r, j, step)
		}
	}
}
