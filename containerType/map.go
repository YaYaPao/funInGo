package main

import (
	"fmt"
)

func defineMap() {
	// 通过 make 来声明一个 map，且初始化的 m 不等于 nil
	// 因为如果 map 为一个 nil，则不能再对其进行赋值
	m := make(map[string]float64)
	m["a"] = 1
	m["b"] = 2
	m["c"] = 3
	fmt.Println(m)
	// 遍历输出
	for k, v := range m {
		fmt.Printf("key is %v and value is %v\n", k, v)
	}
	// 取值，如果没有对应字段，则会返回默认值以及一个该字段是否存在的标识
	value1, hasKey := m["a"]
	fmt.Println(value1, hasKey)
	// value2 is 0 and hasKey is false
	value2, hasKey := m["p"]
	fmt.Println(value2, hasKey)
}

func clearMap() {
	m := make(map[string]string)
	m["a"] = "hello"
	m["b"] = "world"
	fmt.Println(m)
	// delete 方法删除指定 key-value
	delete(m, "a")
	fmt.Println(m)
	m["c"] = "yy"
	// 清空 map
	m = make(map[string]string)
	fmt.Println(m)
}

func main() {
	defineMap()
	clearMap()
}
