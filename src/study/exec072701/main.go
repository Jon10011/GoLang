package main

import (
	"fmt"
	"strings"
)

func makeSuffix(suffix string) func(string) string {
	return func(s string) string {
		if strings.HasSuffix(s, suffix) {
			return s
		}
		return s + suffix
	}
}

func main() {
	funcTest := makeSuffix(".jpg")
	fmt.Println(funcTest("a.jpg"))
	fmt.Println(funcTest("b"))
}
