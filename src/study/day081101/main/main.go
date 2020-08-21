package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name  string
	Age   int
	Score string
	Skill string
}

func testStruct() {
	student := Student{
		Name:  "宋冬",
		Age:   18,
		Score: "100",
		Skill: "篮球",
	}

	student_data, err := json.Marshal(&student)
	if err != nil {
		fmt.Print("报错 err=%v", err)
	}
	fmt.Printf("序列化以前student=%v \n", student)
	fmt.Printf("序列化以后student=%v \n", string(student_data))
}

func testMap() {
	var a map[string]interface{}
	a = make(map[string]interface{})
	a["name"] = "张焕"
	a["age"] = 2
	a["address"] = "合肥"

	a_data, err := json.Marshal(a)
	if err != nil {
		fmt.Println("报错：err=%v", err)
	}
	fmt.Println("a序列化以后是", string(a_data))
}

func testSlice() {
	var slic []map[string]interface{}
	var map1 map[string]interface{} = make(map[string]interface{})
	map1["name"] = "jack"
	map1["age"] = 14
	map1["地址"] = "南京"
	slic = append(slic, map1)

	var map2 map[string]interface{} = make(map[string]interface{})
	map2["name"] = "james"
	map2["age"] = 16
	map2["地址"] = "上海"
	slic = append(slic, map2)

	b_data, err := json.Marshal(slic)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b_data))
}

func main() {
	testStruct()
	testMap()
	testSlice()
}
