package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 2000)
	exitChan := make(chan bool, 4)
	start := time.Now().Unix()

	//开启一个协程，向intchan放入1-8000个数字
	go func(c chan int) {
		for i := 1; i <= 80000; i++ {
			c <- i
		}
		close(c)
	}(intChan)
	//开启4个协程，从intchan中取出数据，判断是否位素数如果是就放入primeChan
	for i := 0; i < 4; i++ {
		go func(c chan int, p chan int, e chan bool, id int) {
			var flag bool
			for {
				num, ok := <-c
				if !ok {
					break
				}
				flag = true
				for i := 2; i < num; i++ {
					if num%i == 0 {
						flag = false
						break
					}
				}
				if flag {
					p <- num
				}
			}
			exitChan <- true
		}(intChan, primeChan, exitChan, i)
	}
	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
			fmt.Println("协程" + strconv.Itoa(i) + "退出了")
		}
		end := time.Now().Unix()
		fmt.Println("===>>>>耗时:", end-start)
		//当从exitChan中取出了4个结果，就可以关闭primeChan管道了
		close(primeChan)
		}()
	//主线程
	for {
		_, ok := <-primeChan
		if !ok {
			break
		}
		// fmt.Println("素数=", res)
	}
	fmt.Println("main主线程退出")

}
