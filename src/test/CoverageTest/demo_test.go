package demo

import "testing"

func TestArea(t *testing.T) {
	area := area(40, 50)
	if area != 2000 {
		t.Error("测试失败")
	}
}

func BenchmarkArea(b *testing.B) {
	for i := 0; i < b.N; i++ {
		area(40, 50)
	}
}
