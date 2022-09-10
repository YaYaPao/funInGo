package main

import "fmt"

// go 内没有 class 的概念，我们可以将函数绑定到 struct 上，来作为方法使用
// method 声明格式：func (v struct) name type {}

type Person struct {
	Age, Height int
}

// GetAge 这里为 `Person` struct 挂载 `GetAge` 方法
// 同时 `GetAge` 也可以作为普通函数调用，或者说其本质就是一个函数，给 structs 充当 methods 而已
func (v Person) GetAge() int {
	return v.Age
}

func main() {
	p := Person{10, 12}
	fmt.Println(p.GetAge())
}
