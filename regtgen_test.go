package regtgen

import (
	"testing"
)

func BenchmarkGeneration(b *testing.B) {
	ch2 := []byte("abcdefghijklmnopqrstuvwxyz")
	exp := Init("###-###", ch2)
	for i := 0; i < b.N; i++ {
		_, _ = exp.Generate()
	}
}
