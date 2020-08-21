package main

import "fmt"

func fbn(n int) []uint64 {
	// slic := make([]uint64, n)
	var slic []uint64 = make([]uint64, n)
	slic[0] = 0
	slic[1] = 1
	for i := 2; i < n; i++ {
		slic[i] = slic[i-1] + slic[i-2]
	}
	
	return slic
}

func main() {
	slic := fbn(100)
	fmt.Println(slic)
}
