package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func writeDataToFile(file *os.File, a int) {
	file.WriteString(strconv.Itoa(a) + "\n")
}

func readFile(file string) []int {
	bytes, _ := ioutil.ReadFile(file)
	str := strings.Split(string(bytes), "\n")
	var result []int = make([]int, 0)
	fmt.Printf("result:%v\n", str)
	for _, v := range str {
		v, _ := strconv.Atoi(v)
		result = append(result, v)
	}
	return result
}

type SortBy []int

func (a SortBy) Len() int           { return len(a) }
func (a SortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool { return a[i] < a[j] }

func main() {

	filePath1 := "./test.txt"
	// filePath2 := "./test2.txt"
	file, _ := os.Create(filePath1)
	rand.Seed(time.Now().Unix())
	for i := 1; i <= 1000; i++ {
		n := rand.Intn(1000000)
		go writeDataToFile(file, n)
	}
	// var intChan chan int
	str := readFile(filePath1)
	sort.Sort(SortBy(str))
	fmt.Println(str)
	for k, v := range str {
		fmt.Printf("k=%v,v=%v\n", k, v)
	}

}
