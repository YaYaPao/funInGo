package main

import (
	"fmt"
)

func add(a *int) {
	*a++
	// now pointer's address has changed and point to new address
	fmt.Printf("current value is %v and it's address is %v\n", *a, &a)
}

func main() {
	a := 1
	fmt.Printf("cache address is %v \n", &a)
	// define a pointer
	var p *int = &a
	// 通过 * 对指针变量进行取值
	fmt.Printf("a's pointer is %v \np's value is %v\n", p, *p)

	*p++
	fmt.Printf("p's value has changed to %v and it's address is %v\n", *p, p)

	add(p)
	// 可以发现，p 的指针地址没变，但是其值发生了改变
	fmt.Printf("p's value has changed to %v and it's address is %v\n", *p, p)

	// 通过 nil 判断空指针
	var p1 *int
	if p1 == nil {
		fmt.Println("p1 is nil pointer")
	}
}
