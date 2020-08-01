package main

import "fmt"

func formatter () {
	s := "hello"
	/*
	通用打印类型
	%v 表示当前变量值
	%T 表示 Go 内的值类型
	 */
	fmt.Printf("type is %T \n", s)
	fmt.Printf("value is %v \n", s)
	/*
	bool 类型
	%t 返回 true or false
	 */
	var b bool
	fmt.Printf("bool is %t \n", b)
	/**
	int 类型
	%d 十进制
	%c 对应的 Unicode
	%b 二进制
	 */
	fmt.Printf("decimal is %d \n", 7)
	fmt.Printf("binary is %b \n", 7)
	/**
	float 类型
	%.2f 保留 2 位小数
	%f 默认保留 6 位小数
	 */
	fmt.Printf("reserve 2 decimal is %.2f \n", 1.23456)
	fmt.Printf("reserve 2 decimal is %f \n", 1.23456)
}

func main() {
	formatter()
}
