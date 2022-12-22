package main

import "fmt"

// 在 go 内，通过 [startIndex:endIndex] 来对数组进行 slice 操作（和 Python 类似）
// 前闭后开
// slice 更像是描述底层数组的某部分，改变切片的值会同步改变数组的值
// 通过理解 slice 的 length 和 capacity 来理解切片和底层数组
// - len(s) 获取切片当前所含的元素个数
// - cap(s) 获取切片底层数组所含的元素个数
func sliceArray() {
	// nums 是数组(array)
	nums := [4]int{1, 2, 3, 4}
	// newNum 是切片(slice)
	var newNum []int = nums[0:3]
	fmt.Println("nums is", nums)
	fmt.Println("newNum is", newNum)

	newNum[0] = 7
	fmt.Println("nums is changed to", nums)
	fmt.Println("newNum is changed to", newNum)
}

// 通过 nil 来判断切片是否为空，其 length 和 capacity 均为 0
// 通过 `append` 向其添加元素，`func append(s []T, vs ...T) []T`
// ⚠️ 注意需要将将返回重新赋值给原切片（与 javascript 行为不同），可以同时添加多个元素
func emptySlice() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("s is nil!")
	}
	// append
	s = append(s, 1)
	s = append(s, 2, 3)
	fmt.Println(s)
}

// 通过 `range` 方法来遍历 slice，每次循环会返回 index 和 value（会复制当前值）
func loopSlice() {
	s := []int{1, 2, 3, 4}
	for i, v := range s {
		fmt.Println(i, v)
	}
}

/*
appendSlice 在 go 内，通过 `append` 合并元素/slice
*/
func appendSlice() {
	var s []string
	// 添加元素
	s = append(s, "hello", "world")
	// 合并 slice，注意在 slice 后添加 ...
	s = append(s, []string{"tony", "young"}...)
}

func main() {
	//sliceArray()
	//emptySlice()
	loopSlice()
}
