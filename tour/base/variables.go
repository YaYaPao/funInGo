package main

import "fmt"

// 通过 `var` 在函数外声明一组变量，并指定类型
// 默认值：int 默认值为 0，string 默认值为 ""，bool 默认值为 false
var a, b int

// 同时，可以为变量赋默认值，此时变量类型可以省略，go 最自动推断其类型
var c = 3

// 通过 `const` 声明一个常量
const Pi = 3.14

// `:=`（short variable declarations） 仅支持在函数体内
func sum() int {
	sum := a + b + c
	return sum
}

// 通过 `T(v) 将 v 转换成 T 类型`
// ⚠️ 整数转字符串，会根据 ASCII 字符集进行对应，而不是将 100 转换为 "100"，因此 go 内会建议先转为 rune（int32，代表 Unicode）类型
func transform() (float32, string) {
	var t1 int = 100
	t2 := float32(t1)
	t3 := string(rune(t1))
	fmt.Printf("t2 is %v, t3 is %s\n", t2, t3)
	fmt.Printf("t2 type id %T, t3 type is %T", t2, t3)
	return t2, t3
}

func main() {
	transform()
}
