package main

import (
	"fmt"
)

//闭包
//累加器
func addUpper() func(int) int {
	var n int = 10
	var str = "hello"
	return func(x int) int {
		n = n + x
		str += string(x)
		fmt.Println("str=", str)
		return n
	}
}

func main() {
	f := addUpper()
	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))
}
