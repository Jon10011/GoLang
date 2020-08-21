package main

import (
	"fmt"
)

func main() {
	var slice []string = make([]string, 5, 5)
	slice[0] = "nihao"
	fmt.Printf("%p\n", &slice)
	slice[1] = "nihaoa"
	fmt.Printf("%p\n", &slice)
	slice[3] = "你好golang"
	fmt.Printf("%p\n", &slice)

	fmt.Println(slice)

	var sli []int = make([]int, 5)
	sli = append(sli, 1)
	sli = append(sli, 2)
	sli = append(sli, 3)
	sli = append(sli, 4)
	fmt.Printf("%p\n", &sli)
	sli = append(sli, 5)
	fmt.Printf("%p\n", &sli)
	sli = append(sli, 6)
	fmt.Printf("%p\n", &sli)
	fmt.Printf("sli=%v \n", sli)

	var sli1 []int = make([]int, 0, 2)
	sli1 = append(sli1, 9)
	fmt.Printf("%d,%d, %p\n", len(sli1), cap(sli1), sli1)
	sli1 = append(sli1, 8)
	fmt.Printf("%d,%d, %p\n", len(sli1), cap(sli1), sli1)
	sli1 = append(sli1, 7)
	fmt.Printf("%d,%d, %p\n", len(sli1), cap(sli1), sli1)
	fmt.Printf("sli1=%v \n", sli1)

	b := make([]int, 0, 2)
	fmt.Printf("%d, %d, %p\n", len(b), cap(b), b) // 0, 2, 0xc420012190
	b = append(b, 1)
	b = append(b, 2)
	fmt.Printf("%d, %d, %p\n", len(b), cap(b), b) // 2, 2, 0xc420012190
	b = append(b, 3)
	fmt.Printf("%d, %d, %p\n", len(b), cap(b), b) // 3, 4, 0xc420012198

}
