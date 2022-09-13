package main

import "fmt"

// go 内没有 class 的概念，我们可以将函数绑定到 struct 上，来作为方法 method 使用
// method 本质是一个绑定了指定 receiver 的函数（with a receiver arguments ）
// method 声明格式：func (v receiver) name type {}
//  - 这意味着不必是 struct，只要是在当前 package 定义的类型均可以为其绑定 method
// 在 Go 内，method 同时支持 value receiver 和 pointer receiver（👌推荐），在 package 内，一般仅允许选择一种 receiver
//  - pointer receiver 更强大，能够直接改变 receiver 指向的值
//  - pointer receiver 更高效，不必进行值的拷贝，减少了性能损耗

// MyInt 定义一个 MyInt 类型，并为其绑定 `SomeIntTest` 方法
type MyInt int

func (v MyInt) SomeIntTest() int {
	if v > 0 {
		return int(v + 1)
	}
	return int(v - 1)
}

// Person 定义一个 Person 的结构体类型，并分别测试其 value receiver 和 pointer receiver
type Person struct {
	Age, Height int
}

// GetAge 这里为 `Person` struct 挂载 `GetAge` 方法
// Value receiver 传递值的拷贝
// 同时 `GetAge` 也可以作为普通函数调用，或者说其本质就是一个函数，给 structs 充当 methods 而已
func (v Person) GetAge() int {
	return v.Age
}

// SetAge SeAge 这里是将原值的指针对象传递作为 receiver
// Pointer receiver
// 这里传递的是原值，而不是值的拷贝
func (v *Person) SetAge(h int) {
	v.Age = h
}

func main() {
	// MyInt
	fmt.Println(MyInt(10).SomeIntTest())
	// Person struct
	p := Person{10, 12}
	fmt.Println(p.GetAge())
	// ⚠️注意，这里我们直接传递值而不是地址，这是因为 Go 解析到 `SetAge` 接受一个指针，因此在内部做了转换
	// 因此也可以使用 `&Person` 来调用 method
	p.SetAge(80)
	fmt.Println(p)
	p2 := &Person{10, 12}
	p2.SetAge(66)
	fmt.Println(p2.GetAge())

}
