package task33102

import "testing"

func BenchmarkSomeMath(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SomeMath(float64(i))
	}
}
