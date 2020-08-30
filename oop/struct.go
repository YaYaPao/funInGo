package main

import (
	"fmt"
)

/**
define a Person struct
name should be Uppercase
type and struct is keyword
Struct is Value Type
*/
type Person struct {
	name string
	age  int
}

func generate() {
	// 1. 通过 var 进行实例化，其初始值为各个类型的缺省值
	var p1 Person
	fmt.Println("before set value p1 is", p1)
	fmt.Printf("init p1: value is %v, type is %T\n", p1, p1)
	// p1 实例化之后，为其进行赋值
	p1.age = 27
	p1.name = "hello"
	fmt.Println("after set value p1 is ", p1)

	// 2. 结合海象表达式是实例化
	p2 := Person{
		name: "david",
		age:  28,
	}
	fmt.Println("p2's value is", p2)

	// 3. 简单声明格式，不通过属性名来对应，直接通过赋值顺序进行控制
	p3 := Person{"lucy", 18}
	fmt.Println("p3's value is", p3)
	fmt.Printf("p3's name is %v and age is %v\n", p3.name, p3.age)
}

// new() 相当于语法糖，对程序语言本身没有影响，但是更方便程序员使用
func sugarInStruct() {
	fmt.Println("------- new a struct --------")
	// 使用内置函数 new() 对结构体进行实例化，形成指针类型的结构体，new() 内置函数会分配内存
	// 通过 new() 实例化之后返回的是该类型新分配的值的**指针**
	p := new(Person)
	// 可以观察到 value 为 &{ 0}，相当于初始化一个 Person 实例并返回其地址
	fmt.Printf("test new(): value is %v and type is %T and address is %p\n", p, p, p)
	// 本来需要通过 `(*p).age` 从地址取出值之后，对其进行赋值，但是由于语法糖，可以让我们直接对其进行赋值
	// 这样的语法糖同样存在于数组和切片中
	p.age = 18
	p.name = "dav"
	fmt.Printf("value is %v\n", p)
}

// deep clone or simple clone
func clone() {
	// 值类型是深拷贝，就新的对象重新分配内存
	s1 := Person{"test1", 1}
	s2 := s1
	s2.name = "test2"
	// 可以发现 s1 的值并没有发生改变，复制已经完成了深拷贝
	fmt.Println("s1 is", s1)
	fmt.Println("s2 is", s2)

	s3 := &s1
	s3.name = "test3"
	// 将 s1 的地址进行传递，同时结合语法糖对其进行重新赋值，可以发现 s1 随着 s3 的赋值而改变，即为浅拷贝
	// 同理，通过 new() 进行实例化之后进行复制，实际上也是传递的指针，也是浅拷贝
	// s2 并没有发生改变，印证了 s2 值在新的一块内存中
	fmt.Println("s1 is", s1)
	fmt.Println("s2 is", s2)
	fmt.Println("s3 is", s3)
}

func shadowStruct() {
	println("------ shadow struct define ------")
	// 这里声明一个匿名结构体，并将其赋值给 s
	s := struct {
		name string
		sex  bool
	}{"lucy", false}
	fmt.Println(s)
	// 结构体的匿名字段表示的是在结构体内没有名字，只包含一个没有字段名的类型
	s2 := struct {
		int8
		string
	}{1, "hello shadow"}
	fmt.Println(s2)
}

func combine() {
	// 聚合就是声明一个新的结构体，同时指定字段对应已存在的结构体
	// 一个类是另一个类的属性，访问时需要通过 `实例.属性.属性` 来进行访问
	// 可以发现，实际上对聚合的结构体进行了一次深拷贝，因此两者互不干扰
	s := struct {
		person Person
		sex    bool
	}{}
	p := Person{"luck", 18}
	s.person = p
	s.sex = false
	fmt.Println(s)
	p.name = "david"
	fmt.Println("after p changed name ", s, p)
	// 注意访问方式
	s.person.age = 100
	fmt.Println("after s changed person.age", s, p)
}

func inherit() {
	// 一个类作为另一个类的子类，注意理解和聚合的区别
	fmt.Println("------ inherit ------")
	s := struct {
		Person
		sex bool
	}{}
	// 通过匿名结构体字段来实现继承，其赋值方式已经发生改变
	s.age = 1
	s.name = "uu"
	s.sex = true
	fmt.Println(s)
}

func main() {
	// 实例化
	generate()
	// new() 语法糖
	sugarInStruct()
	// 深/浅拷贝
	clone()
	// 匿名结构体和匿名结构体字段
	shadowStruct()
	// 实现聚合和继承
	combine()
	inherit()
}
