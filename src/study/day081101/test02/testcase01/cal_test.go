package main

import "testing"

func TestAddUpper(t *testing.T)  {
	res := addUpper(20)
	if res!=55{
		t.Fatalf("addUpper(10)执行错误，期望值=%v，实际值=%v\n",55,res)
	}
	t.Logf("addUpper(10)执行正确。。。。。。")
}