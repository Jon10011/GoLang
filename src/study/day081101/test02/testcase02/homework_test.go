package main

import (
	"testing"
)

func TestHomework(t *testing.T) {
	filePath := "./demo.txt"
	monster := Monster{
		Name:  "宋冬",
		Age:   19,
		Skill: "篮球",
	}
	res := monster.Store(filePath)
	if !res {
		t.Fatalf("monster.store()错误，希望值为%v，实际为%v", true, res)
	}
	t.Logf("monster.store()执行成功")
}

func TestHomeworkRestore(t *testing.T) {
	filePath := "./demo.txt"
	var monster Monster
	res := monster.ReStore(filePath, "宋冬")
	if !res {
		t.Fatalf("monster.restore()错误，希望值为%v，实际为%v", "宋冬", monster.Name)
	}
	t.Logf("monster.restore()执行成功")
}
