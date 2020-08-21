package main

import "fmt"

func monkeyP(n int) int {
	if n > 10 || n < 1 {
		return 0
	}
	if n == 10 {
		return 1
	}
	return (monkeyP(n+1) + 1) * 2
}

func main() {
	peachNum := monkeyP(10)
	fmt.Printf("桃子的数量是 %v", peachNum)
}
