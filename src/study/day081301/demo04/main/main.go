package main

import "fmt"

//channel的使用

func main() {
	var inChan chan int
	inChan = make(chan int, 3)
	fmt.Println(inChan)

	inChan <- 10
	fmt.Println(cap(inChan), len(inChan))

	var num int
	num = <-inChan
	fmt.Println(num)
	fmt.Println(cap(inChan), len(inChan))

}
