package main

import "fmt"

// go 内支持指针，指针指向内存地址
// 在 go 内：
// - 通过 `&` 对变量进行取址操作，即取地址
// - 通过 `*` 对内存地址进行取值操作，即取地址对应的值
// 如果声明变量还未赋值，则不会为其分配内存地址，此时取址为空
func main() {
	data := "123"
	p := &data
	fmt.Printf("data address is %v\n", p)
	*p = "456"
	fmt.Println("Now data has changed to", data)

	var v *string
	if v == nil {
		fmt.Printf("v is a null point")
	}
}
