package main

import (
	"fmt"
	"math"
)

var f = func(num float64) {
	// 求平方根
	v := math.Sqrt(num)
	fmt.Println(v)
}

func main() {
	// 自执行
	func(str string) {
		fmt.Println(str)
	}("hello world")
	// 传递给参数
	f(100)
}
