package main

import "fmt"

func test1() {
	source := 20
	if source % 2 == 0 {
		fmt.Printf("%v is 偶数 \n", source)
	} else {
		fmt.Printf("%v is 奇数", source)
	}

	// 在条件语句内声明变量
	if a := 21; a % 2 == 1 {
		fmt.Printf("%v is male", a)
	} else {
		fmt.Printf("%v is female", a)
	}
}

func main() {
	test1()
}
