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

func main() {
	arrayBase()
	nums := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	twoDimensional(nums)
	// verify the nums has change or not
	fmt.Println("after changed nums is", nums)
}
