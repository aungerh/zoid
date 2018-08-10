package regtgen

import (
	"testing"
)

func BenchmarkGeneration(b *testing.B) {
	ch2 := []byte("abcdefghijklmnopqrstuvwxyz")
	exp := InitGenerator("##-##-##-##", ch2)
	_, _ = exp.Generate(1000)
}
