package main

import (
	"fmt"
	"runtime"
)

//MPG模式

func main() {
	num := runtime.NumCPU()
	num2 := runtime.NumGoroutine()
	runtime.GOMAXPROCS(num - 1)
	fmt.Println(num)
	fmt.Println(num2)

}
