package main

import (
	"fmt"
	"study/day080503/model"
)

//面向对象三大特性：封装、继承、多态

//封装：将共有的属性抽象出来；好处：隐藏实现的细节，结构体的方法对属性的验证
//继承：解决代码冗余，使用匿名结构体来实现继承的特性

func main() {
	p := model.NewPerson("songdong")
	fmt.Println(p)
	p.SetAge(19)
	p.SetSal(5000)
	fmt.Println(p.GetAge())
	fmt.Println(p.GetSal())

	account := model.NewAccount("songdong")
	account.SetPwd("123456")
	fmt.Println(account)
}
