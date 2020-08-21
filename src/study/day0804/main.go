package main

import "fmt"

//定义结构体
type Cat struct {
	Name  string
	Age   int
	Color string
}

type Person struct {
	Name   string
	Age    int
	Scores [5]float64
	ptr    *int
	slice  []int
	map1   map[string]string
}

type Student struct {
	Name string
	Age  int
}

func main() {
	var cat1 Cat
	cat1.Name = "小白"
	cat1.Age = 3
	cat1.Color = "白色"
	fmt.Println(cat1)
	fmt.Println("--------------------------")
	var person Person
	fmt.Println("person=", person)
	//使用slice
	person.slice = make([]int, 10)
	person.slice[0] = 100
	fmt.Println("person=", person)

	//使用map
	person.map1 = make(map[string]string)
	person.map1["name"] = "songdong"
	person.map1["age"] = "18"
	fmt.Println("person=", person)

	//创建struct的几种方式
	//1
	var student1 Student
	student1.Name = "songdong"
	student1.Age = 18
	fmt.Println("student1=", student1)
	//2
	student2 := Student{"Jon", 18}
	fmt.Println("student2=", student2)
	//3
	var student3 *Student = new(Student)
	(*student3).Name = "Jack"
	student3.Name = "jim"

	(*student3).Age = 12
	student3.Age = 16
	fmt.Println("student3=", student3)

	//4,因为student4是一个指针，因此标准的访问字段的方法是(*student4).Name="james"
	var student4 *Student = &Student{}
	// var student4 *Student = &Student{"tom",14}
	(*student4).Name = "Jack4"
	student4.Name = "jim4"

	(*student4).Age = 124
	student4.Age = 164
	fmt.Println("student4=", student4)

}
