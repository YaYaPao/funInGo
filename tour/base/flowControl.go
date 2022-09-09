package main

import (
	"fmt"
	"time"
)

// go 中仅支持 `for` 一种循环体，它由三部分组成：
// 1. 初始化表达式，在循环开始前执行
// 2. 条件表达式，在每次循环开始前判断
// 3. post statement, 在每次循环结束后执行
// 其中，初始化变量（`i`）仅能够在 for 循环内访问到
func loopSum(data int) int {
	sum := 0
	for i := 0; i < data; i++ {
		sum += i
	}
	return sum
}

// for 循环内 `init` 和 `post statement` 表达式可以省略，此时类似 `while` 用法
func loopWithoutInit() {
	sum := 1
	for sum < 1000 {
		fmt.Println("before:", sum)
		sum += sum
		fmt.Println("after:", sum)
	}
	fmt.Println(sum)
}

// 在 go 中，`if` 和 `for` 类似，条件表达式都不需要 () 进行包裹
// 同样，也支持在判断前对值进行计算
// 条件体内声明的变量只能在条件作用域内（包括 else 结构体内）进行访问
func conditionCaseIf(data int) bool {
	if v := data * 2; v > 0 {
		return true
	} else {
		return false
	}
}

// 与其他语言的区别在于：
// 1. 在 go 内，不需要使用 `break` 语句，它会找到命中的语句执行后，立即调出循环；
// 2. case 语句支持计算；
func conditionSwitch() {
	fmt.Println("When's Sunday?")
	today := time.Now().Weekday()
	switch time.Sunday {
	case today + 0:
		fmt.Println("Today is.")
	case today + 1:
		fmt.Println("Tomorrow is.")
	default:
		fmt.Println("Far than 2 days.")
	}
}

// 对于长 if-then-else 语句，为了降低代码的圈复杂度，我们可以用 switch 语句进行改写
func conditionIfToSwitch(data int) {
	switch {
	case data > 10:
		fmt.Println("Number is greater than 10.")
	case data > 5:
		fmt.Println("Number is greater than 5.")
	default:
		fmt.Println("Number is less than 5.")
	}
}

// `Defer` 是 go 中的一个新概念，类似 JavaScript 的 setTimeout，但是又有些许区别
// defer 指定的步骤会在整个函数执行完毕之后依次执行
// defer 标记的方法会被写入一个堆栈内，遵循先进后出的原则
func conditionDefer(data int) {
	for i := 0; i < data; i++ {
		defer fmt.Println(i)
	}
	fmt.Printf("Print the i\n")
}

func main() {
	//fmt.Println(loopSum(10))
	//conditionSwitch()
	conditionDefer(10)
}
