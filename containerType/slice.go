package main

import (
	"fmt"
)

func displayBase() {
	// capacity 表示可选参数，初始化一个空的切片
	slice1 := make([]int, 3, 5)
	fmt.Printf("slice1'type is %T\n", slice1)
	fmt.Printf("len = %d and cap is %d and value = %v\n", len(slice1), cap(slice1), slice1)

	// 直接初始化一个带值切片
	slice2 := []int{1, 2, 3}
	fmt.Printf("type is %T and value is %v\n", slice2, slice2)

	// 通过数组截取来初始化切片
	arr := [5]int{1, 2, 3, 4, 5}
	// 包前不包后
	s := arr[1:3]
	fmt.Println(s)
}

func testType() {
	arr := [5]int{1, 2, 3, 4, 5}
	s1 := arr[:]
	s2 := arr[0:3]
	s3 := s1[0:3]
	fmt.Println(s1, s2, s3)
	// 三个切片的地址不同，但是底层都引用同一个数组
	fmt.Printf("%p, %p, %p", &s1, &s2, &s3)
	// 可以观察到 s1, s2, s3 的第一个元素都被更改为 77
	s1[0] = 77
	fmt.Println(s1, s2, s3)
}

func testMethod() {
	arr := [3]int{1, 2, 3}
	s1 := arr[:]
	s2 := []int{7, 8}
	/**
	append() 用于向切片内追加一个或者多个元素，也可以追加一个切片
	append() 会影响切片的底层数组
	当容量不够时，即 (cap - len == 0) 时，Go 会创建一个新的内存地址来存储元素
	*/
	s1 = append(s1, 4, 5)
	fmt.Println(s1, arr)
	// append 追加一个切片
	s2 = append(s2, s1...)
	fmt.Println("s2 is", s2)

	// 可以发现 arr 内的元素并没有被更改
	s1[2] = 77
	fmt.Println(s1, arr)

	// copy() 复制一个切片
	s3 := make([]int, 5)
	count := copy(s3, s1)
	fmt.Println("copy number is", count, s3)
}

func main() {
	displayBase()
	testType()
	testMethod()
}
