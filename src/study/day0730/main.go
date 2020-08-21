package main

import "fmt"

func main() {
	//数组与数组遍历
	var lis [8]int
	lis[0] = 1
	lis[2] = 2
	lis[3] = 3
	lis[4] = 4
	lis[5] = 5
	lis[6] = 6
	fmt.Println(lis)
	for _, value := range lis {
		fmt.Printf("value=%v\n", value)
	}

}
