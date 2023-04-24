package utils

import "testing"

func TestAddUpper(t *testing.T) {
	//调用
	res := AddUpper(10)
	if res != 55 {
		t.Fatalf("错误")
	}

	t.Logf("正确")
}
