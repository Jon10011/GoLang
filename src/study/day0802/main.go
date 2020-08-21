package main

import (
	"fmt"
)

func main() {
	str := "hello宋冬"
	arr1 := []byte(str)
	arr1[5] = 'z'
	str = string(arr1)
	fmt.Println("arr1=", str)

	arr2 := []rune(str)
	arr2[0] = '你'
	str2 := string(arr2)
	fmt.Println("arr2=", str2)
}
