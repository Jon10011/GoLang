package main

import "fmt"

type Student struct {
	name   string
	gender string
	age    int
	id     int
	score  float64
}

type Box struct {
	len    float64
	width  float64
	height float64
}

type Vistor struct {
	Name string
	Age  int
}

func (vistor *Vistor) showPrice() {
	if vistor.Age > 18 {
		fmt.Printf("游客的名字为%v 年龄为%v 收费20元 \n", vistor.Name, vistor.Age)
	} else {
		fmt.Printf("游客的名字为%v 年龄为%v 免费 \n", vistor.Name, vistor.Age)
	}
}

func (box *Box) getVolum() float64 {
	return box.len * box.width * box.height
}

func (student *Student) say() string {
	infoStr := fmt.Sprintf("student的信息 name=[%v] gender=[%v] age=[%v] id=[%v] score=[%v]", student.name, student.gender, student.age, student.id, student.score)
	return infoStr
}

func main() {
	var student1 = Student{
		name:   "jon",
		gender: "men",
		age:    18,
		id:     1,
		score:  99.99,
	}
	fmt.Println(student1.say())

	var box = Box{
		1.2,
		2.5,
		9.9,
	}
	fmt.Printf("volum=%.2f \n", box.getVolum())

	var vistor Vistor
	for {
		fmt.Println("请输入你的名字：")
		fmt.Scanln(&vistor.Name)
		if vistor.Name == "exit" {
			fmt.Println("退出程序。。。。。")
			break
		}
		fmt.Println("请输入你的年龄：")
		fmt.Scanln(&vistor.Age)
		vistor.showPrice()
	}
}
