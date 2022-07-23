package benchmarks

import (
	"advanced-io-operations/impl"
	"fmt"
	"math/rand"
	"strings"
	"testing"
)

func BenchmarkScrambleWriter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s strings.Builder
		w := impl.NewScrambleWriter(&s, rand.New(rand.NewSource(1)), 0.5)
		fmt.Fprint(w, "Hello! this is a sample text. \nCan you read it? Yes")
		fmt.Println(s.String())
	}
}
