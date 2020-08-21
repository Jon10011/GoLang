package main

import "fmt"

func worker(c chan int) {
	fmt.Println("i am a worker")
	num := <-c
	fmt.Println(num)
}

func main() {
	c := make(chan int)
	go worker(c)
	c <- 2
	fmt.Println("i am a main")
}
