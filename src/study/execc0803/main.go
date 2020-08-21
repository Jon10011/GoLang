package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "hellosongdong123nihaoa.jpg"
	fmt.Println(strings.LastIndex(str1, "."))
	slic := str1[strings.LastIndex(str1, "."):]
	fmt.Println("slic=", slic)

	slic1 := []rune(str1)
	slic1[0] = '你'
	fmt.Println("slic1=", slic1)
	str1 = string(slic1)
	fmt.Println("str1=", str1)
}
