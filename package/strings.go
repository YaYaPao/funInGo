package main

import (
	"fmt"
	"strings"
)

func searchString() {
	a := "hello"
	b := "world"
	// 从词典上比较两个字符串
	com := strings.Compare(a, b)
	fmt.Println("compare result is ---", com)
	hasValue := strings.Contains(a, "e")
	fmt.Printf("%v has e is %v\n", a, hasValue)
	count := strings.Count(a, "l")
	fmt.Printf("%v has %v l\n", a, count)
}

func splitString() {
	s := "hello world!"
	fmt.Println(s)
}

func main() {
	// test strings methods
	// you can link https://golang.org/pkg/strings/#Compare for more details
	// just have fun!
	searchString()
	splitString()
}
