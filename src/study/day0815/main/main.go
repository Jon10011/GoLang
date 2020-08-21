package main

import (
	"fmt"
)

func writeData(intChan chan int) {
	for i := 1; i <= 50; i++ {
		intChan <- i
		fmt.Printf("writeData %v \n", i)
		// time.Sleep(time.Second)
	}
	close(intChan)
}

func readData(boolChan chan bool, intChan chan int) {
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Printf("readData读取到数据%v\n", v)
		// time.Sleep(time.Second)
	}
	//readData读取完以后任务完成
	boolChan <- true
	close(boolChan)

}

func main() {
	intChan := make(chan int, 50)
	boolChan := make(chan bool, 1)

	go writeData(intChan)
	go readData(boolChan, intChan)

	// time.Sleep(time.Second * 10)
	for {
		_, ok := <-boolChan
		if !ok {
			break
		}
	}
}
