package demo

import "testing"

func BenchmarkName(t *testing.B) {
	for i := 0; i < t.N; i++ {
		area(40, 50)
	}
}
