package main

import "fmt"

func main() {
	slic1 := make([]interface{}, 5)
	// var i interface{}
	slic1[0] = "songdong"
	slic1[1] = 18
	slic1[2] = []int{1, 2, 3, 4, 5}
	slic1[3] = []string{"测试1", "测试2"}
	fmt.Println(slic1)

	var any []interface{}
	any = append(any, "songdong")
	any = append(any, 19)
	any = append(any, []int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	any = append(any, []string{"宋冬", "张焕"})
	fmt.Println(any)

}
