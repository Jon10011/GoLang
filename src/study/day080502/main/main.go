package main

import (
	"GoLang_Code/project_demo/chapter10/factory/model"
	"fmt"
)

func main() {
	// var stu = model.Student{
	// 	Name:  "tom",
	// 	Score: 79.5,
	// }

	var stu = model.NewStudent("songdong", 98.8)
	fmt.Println(stu)
	fmt.Println("name=", stu.Name, "score=", stu.GetScore())
}
