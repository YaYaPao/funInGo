package main

import "fmt"

func ghost() (int, int) {
	return 10, 20
}

func setValue () {
	// 声明两个变量
	a, b := 1, 2
	fmt.Println(a, b)
	// 变量多重赋值，从左到右依次进行赋值
	b, a = a, b
	fmt.Println(a, b)
}

func testVariable () {
	// 通过 var 变量名 数据类型 = 初始值
	var a int = 7
	// 可以忽略数据类型，交给编译器进行推断
	var b = "hello"
	// 海象表达式，变量名 := 表达式，首选方式
	c := true
	// 利用海象表达式同时声明多个变量
	c1, c2 := "hello", "world"
	var (
		d int
		e string
		f func()
	)
	fmt.Println(a, b, c, d, e, f)
	fmt.Println(c1, c2)
}

func main() {
	testVariable()
	setValue()
	a, _ := ghost()
	_, b := ghost()
	fmt.Println(a, b)
}
