package main

import (
	"encoding/json"
	"fmt"
)

// 序列化以前student={宋冬 18 100 篮球}
// 序列化以后student={"Name":"宋冬","Age":18,"Score":"100","Skill":"篮球"}
// a序列化以后是 {"address":"合肥","age":2,"name":"张焕"}
// [{"age":14,"name":"jack","地址":"南京"},{"age":16,"name":"james","地址":"上海"}]

type Student struct {
	Name  string
	Age   int
	Score string
	Skill string
}

func unmarshalStruct() {
	str := "{\"Name\":\"宋冬\",\"Age\":18,\"Score\":\"100\",\"Skill\":\"篮球\"} "
	var student Student
	err := json.Unmarshal([]byte(str), &student)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(student)
}

func unmarshalMap() {
	str := "{\"Name\":\"宋冬\",\"Age\":18,\"Score\":\"100\",\"Skill\":\"篮球\"} "
	var a map[string]interface{}
	err := json.Unmarshal([]byte(str), &a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(a)
}

func main() {
	unmarshalStruct()
	unmarshalMap()
}
