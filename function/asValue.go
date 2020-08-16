package main

import (
	"fmt"
	// strings package 用来处理字符串类型
	"strings"
)

// 通过 type 关键字定义一个函数类型
type vertifyNum func(string) string

// 实现定义的函数类型
func transform(str string) string {
	result := ""
	for i, value := range str {
		if i%2 == 0 {
			// 转换成大写
			result += strings.ToUpper(string(value))
		} else {
			result += strings.ToLower(string(value))
		}
	}
	return result
}

// 作为参数调用
// 这里将一个 type 为 vertifyNum 的函数作为参数传入，在函数体内进行执行
// 打印 f 的类型，可以看出为 main.vertifyNum
func dealString(str string, f vertifyNum) string {
	fmt.Printf("%T and %v \n", f, str)
	return f(str)
}

func main() {
	var value = dealString(string("hello, world"), transform)
	fmt.Println(value)
}
