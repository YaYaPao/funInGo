package main

import "fmt"

// const 用来标记元素为常量，在 go 程序运行时不会被修改

const str string = "123"

// const 可以用于声明枚举
const (
	Name   = "David"
	Female = false
)

// 直接使用常量
func main() {
	fmt.Println(Name)
}
