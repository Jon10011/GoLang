package main

import (
	"fmt"
)

//打印一个10层的金字塔
func goldTower(floor int) {
	for i := 1; i <= floor; i++ {
		//在打印*前先打印空格
		for j := 1; j <= floor-i; j++ {
			fmt.Print(" ")
		}
		//每层打印多少个*
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func testMutil(n int) {

	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v * %v = %v ", i, j, i*j)
		}
		fmt.Println()
	}

}

func main() {
	var floor string
	fmt.Println("请输入层数")
	fmt.Scanln(&floor)
	goldTower(20)

	var numFloor int
	fmt.Println("请输入乘法表层数")
	fmt.Scanln(&numFloor)
	testMutil((numFloor))

}
