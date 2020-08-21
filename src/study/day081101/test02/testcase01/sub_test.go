package main

import "testing"

func TestGetSub1(t *testing.T) {
	res := getSub(20, 30)
	if res != 50 {
		t.Fatalf("getSub(10)执行错误，期望值=%v，实际值=%v\n", 55, res)
	}
	t.Logf("getSub(10)执行正确。。。。。。")
}
