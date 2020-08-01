package main

import "fmt"

func main() {
	a := 100
	b := 100.6
	c := 100.2
	// 整数转字符串，会更具 ASCII 字符集进行对应，而不是将 100 转换为 "100"
	fmt.Printf("type: %T, value: %q \n", string(a), string(a))
	fmt.Printf("type: %T, value: %v \n", int(b), int(b))
	fmt.Printf("type: %T, value: %v \n", int(c), int(c))
}
