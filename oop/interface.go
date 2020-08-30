package main

import (
	"fmt"
)

type Show interface {
	show()
}

type Duck struct {
	name  string
	label string
}

type Cat struct {
	name string
	word string
}

// 实现 Duck 的 show 接口
func (d Duck) show() {
	fmt.Printf("duck named %v say %v\n", d.name, d.label)
}

// 实现 Cat 的 show 接口
func (c Cat) show() {
	fmt.Printf("cat named %v say %v\n", c.name, c.word)
}

func main() {
	// 注意理解 interface 和 method 共同实现多态
	// 声明一个 interface 赋值给变量 i
	// **定义一个接口变量，实际上可以赋值给任意一个实现了该接口的变量实例**
	var i Show
	duck := Duck{"peter", "gagaga"}
	duck.show()
	cat := Cat{"kitty", "miao!"}
	cat.show()
	// 将 duck 赋值给 i
	i = duck
	i.show()
	// 将 cat 赋值给 i
	// 实际上同样是 show 方法，但是不同的结构体对其有不同的实现方式，因此类具有多态的特性
	// 注意：接口类型的对象，不能访问其实现类的实例的属性字段
	i = cat
	cat.show()
}
