package demo

import "testing"

func TestArea(t *testing.T) {
	area := area(40, 50)
	if area != 2000 {
		t.Error("测试失败")
	}
}
