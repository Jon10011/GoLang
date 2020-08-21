package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//len按字节返回
	str := "hello你好"
	fmt.Println(len(str))
	fmt.Println("--------------------------------------")
	//字符串遍历，含有中文的字符串需要用  []rune(str)  转化一下
	r := []rune(str)
	for i := 0; i < len(r); i++ {
		fmt.Printf("字符=%c \n", r[i])
	}
	fmt.Println("--------------------------------------")
	//字符串转整数： n，err := strconv.Atoi("12")
	n, err := strconv.Atoi("1")
	if err != nil {
		fmt.Println("trans err=", err)
	} else {
		fmt.Println("trans ok=", n)
	}
	fmt.Println("--------------------------------------")
	//整数转字符串
	str1 := strconv.Itoa(1234)
	fmt.Printf("str1=%v,str1type=%T \n", str1, str1)
	fmt.Println("--------------------------------------")
	//字符串转[]byte
	var bytes = []byte("hello golang")
	fmt.Println("bytes=", bytes)
	fmt.Println("--------------------------------------")
	//查找字串是否在制定的字符串中
	bools := strings.Contains("seafood", "foo1")
	fmt.Println(bools)
	fmt.Println("--------------------------------------")
	//查找字符串中有几个指定的字串
	num := strings.Count("chasersfsssdf", "s")
	fmt.Printf("字符串中含有%v个s字串\n", num)
	fmt.Println("--------------------------------------")
	//字符串的比较 1）“==”需要区分 大小写；2）strings.Equals(str1,str2)不区分大小写
	a := "abc" == "Abc"
	b := strings.EqualFold("abc", "Abc")
	fmt.Println("a=", a)
	fmt.Println("b=", b)

	fmt.Println("--------------------------------------")
	//字串第一次出现的位置
	index := strings.Index("abcderfd,m", ",")
	fmt.Println("index=", index)
	//字串最后一次出现的位置
	lastIndex := strings.LastIndex("adbadbadbadbadb", "adb")
	fmt.Println("lastIndex=", lastIndex)
	fmt.Println("--------------------------------------")

	//将指定的子串替换成另一个子串strings.Replace("go go hello go","go","golang",n)
	//n可以指定你希望替换几个，如果n=-1表示全部替换
	str3 := strings.Replace("go go hello go", "go", "golang", -1)
	fmt.Println("str3=", str3)

	str4 := strings.ReplaceAll("go go hello go", "go", "golang")
	fmt.Println("str4=", str4)

	fmt.Println("--------------------------------------")
	//切割字符串
	s := strings.Split("hello", "")
	fmt.Println("s=", s)
	fmt.Printf("s=%v s的Type是%T\n", s, s)
	fmt.Println("--------------------------------------")
	//大小写转换
	upperStr := strings.ToUpper("ABcsdD")
	lowerStr := strings.ToLower("ABCDFEfjfk")
	fmt.Println("upperStr=", upperStr)
	fmt.Println("lowerStr=", lowerStr)
	fmt.Println("--------------------------------------")
	//z去掉字符串两边的空格
	trimLeftAndRight := strings.TrimSpace("      heiheihei ")
	fmt.Println(trimLeftAndRight)

	fmt.Println("--------------------------------------")
	fmt.Println("--------------------------------------")
	fmt.Println("--------------------------------------")
	fmt.Println("--------------------------------------")

}
