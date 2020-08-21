package main

import (
	"fmt"
	"sync"
)

//channel管道

//全局变量加锁同步

var (
	mymap = make(map[int]int, 10)
	//声明全局互斥锁
	lock sync.Mutex
)

func test(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	lock.Lock()
	mymap[n] = res
	lock.Unlock()
}
func main() {
	for i := 1; i <= 200; i++ {
		go test(i)
	}
	lock.Lock()
	for i, v := range mymap {
		fmt.Printf("mymap[%d]=%d \n", i, v)
	}
	lock.Unlock()
}
