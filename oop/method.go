package main

import (
	"fmt"
)

type Peach struct {
	price float64
	color string
}

type LittlePeach struct {
	Peach
	name string
}

// define method
// 注意 func(var Struct) methodName() {} 的形式
func (p Peach) printDetails() {
	fmt.Println("---- hello i am a method ----")
	// p 用来赋值给类型为 Peach 的结构体实例
	fmt.Printf("price is %v and color is %v\n", p.price, p.color)
}

func printDetails(p Peach) {
	fmt.Println("---- hello i am a function ----")
	fmt.Printf("price is %v and color is %v\n", p.price, p.color)
}

// 这里对继承于 Peach 的方法进行重写，结构体实例在调用时会按照就近原则进行调用
func (p LittlePeach) printDetails() {
	fmt.Println("----- little peach -----")
	fmt.Printf("i am little peach, price is %v, color is %v and name is %v\n", p.price, p.color, p.name)
}

func methodInherit() {
	s := struct {
		Peach
		weight float64
	}{}
	s.price = 12.34
	s.color = "white"
	s.weight = 1.23
	// 这里继承与 Peach 的结构体 s 可以直接调用 printDetails 方法
	s.printDetails()
}

func methodRewrite() {
	s := LittlePeach{Peach{12.34, "red"}, "little peach"}
	// 可以看到这里直接调用 LittlePeach 的 printDetails 方法，而不再向上冒泡
	s.printDetails()
}

func main() {
	p1 := Peach{12.67, "black"}
	// 这里 p1 作为一个 receiver，调用 printDetails 方法
	p1.printDetails()
	printDetails(p1)

	// 实现方法的继承
	methodInherit()
	methodRewrite()
}
