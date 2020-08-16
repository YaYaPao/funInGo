package main

import (
	"fmt"
)

// 函数作用域内，变量查询顺序同 JavaScript，就近查询
func scope() {
	a := "hello in my scope"
	fmt.Println(a)
}

// 可变参数实现
func varParams(data ...float64) (sum, avg float64, count int) {
	for _, value := range data {
		sum += value
		count++
	}
	avg = sum / float64(count)
	return sum, avg, count
}

func main() {
	scope()
	source := []float64{4, 5, 6}
	fmt.Println(varParams(source...))

	sum, avg, count := varParams(1, 2, 3)
	fmt.Printf("sum is %v \navg is %v \ncount is %d \n", sum, avg, count)
}
