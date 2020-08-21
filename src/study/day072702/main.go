package main

import (
	"fmt"
)

//defer
//按照==>>>先入后出<<<==的方式出栈执行，即先打印res，再打印n2，n1
func sum(n1 int, n2 int) int {
	defer fmt.Println("n1=", n1) //3
	defer fmt.Println("n2=", n2) //2
	n1++
	n2++
	res := n1 + n2
	fmt.Println("res=", res) //1
	return res
}

func main() {
	res2 := sum(10, 4)
	fmt.Println("res2=", res2) //4
}
