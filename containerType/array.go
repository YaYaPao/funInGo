package main

import (
	"fmt"
)

func arrayBase() {
	// 定义一维数组，注意元素个数不能超过数组长度
	nums := [3]int{1, 2, 3}
	fmt.Println(nums)
	source := []int{3, 2, 1}
	fmt.Println(source)
	fmt.Printf("soucre length is %v and second element is %v\n", len(source), source[1])

	// 遍历方式1
	for i := 0; i < len(source); i++ {
		fmt.Printf("current index is %v and value is %v\n", i, source[i])
	}
	// 遍历方式2
	for index, value := range nums {
		fmt.Println(index, value)
	}
}

// 二维数组
func twoDimensional(data [][]int) {
	// define 2 rows and 3 columns array
	fmt.Println(data)
	// value type, then change the value
	data[0][0] = 100
	fmt.Println(data)
}

func sugarInArray() {
	a1 := [3]int{1, 2, 3}
	a2 := &a1
	// 这里对 a2 进行直接赋值操作，得益于 Go 内实现的语法糖
	a2[0] = 100
	fmt.Printf("a1 is %v and a2 is %v\n", a1, a2)
}

func main() {
	arrayBase()
	nums := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	twoDimensional(nums)
	// verify the nums has change or not
	// 数组是引用类型，传递的是该值的指针
	fmt.Println("after changed nums is", nums)
	// 测试数组中的语法糖
	sugarInArray()
}
