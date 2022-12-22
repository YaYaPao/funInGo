package main

import "fmt"

// interface 是 method 签名的集合，是实现多态的关键
//  - 配合 `method` 使用，⚠️注意 method 定义时的 receiver 类型
//  - 遵循 `Duck Test` 即鸭子测试：如果一个动作长得像鸭子，走路像鸭子，叫得也像鸭子，那它就是一只鸭子
//  - `interface` 是 Go 内实现多态的关键一环，它定义对象的行为集合
// 在 Go 内实现对象，需要 `struct + interface + methods`
// interface 在 Go 内是隐式实现（约定式），而不需要 `implement` 等关键字
// interface 在 Go 底层实现可以看作是一个 `(value, type)` 的元组，通过 interface 调用其方法等同于调用 value 的同名方法

type Talk interface {
	whisper(data string)
	getName() string
}

type Duck struct {
	name string
}

type Puppy struct {
	name string
}

func quickDuckInterface() {
	var t Talk
	describe(t)
	d := Duck{"Donald Duck"}
	t = &d
	describe(t)
	t.whisper("i am a duck!")
	name := t.getName()
	fmt.Println("Duck name is", name)
}

// 这里 d 仅声明，没有初始化赋值行为，receiver 接受的是一个指针，因此需要 `*Duck` 来取值
func testNil() {
	var t Talk
	var d *Duck
	t = d
	describe(t)
	t.whisper("hi")
}

// 通过空接口 `interface{}` 来实现类似 javascript 的动态变量
// empty interface 类似 any，能够转变为不同的(值 类型)
// 对于 empty interface，我们可以通过 `v := i.(type)` 来进行类型断言
//  - 类似 javascript 内的 `typeof a === 'string' && a`
func emptyInterface() {
	var v interface{}
	fmt.Printf("value is %v, and type is %T\n", v, v)
	v = 123
	fmt.Printf("value is %v, and type is %T\n", v, v)
	v = "123"
	fmt.Printf("value is %v, and type is %T\n", v, v)

	// 使用类型 type assert 来获取值
	target, ok := v.(string)
	fmt.Println(target, ok)
	target2, ok2 := v.(int)
	fmt.Println(target2, ok2)
}

// 可以通过 switch 语句对 interface 类型进行循环后取值
func switchInterface(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Type of %v is %T!", v, v)
	case string:
		fmt.Printf("Type of %v is %T!", v, v)
	default:
		fmt.Println("Unknown type!")
	}
}

// 在 Go 内，最常用的一个 interface 是 fmt 模块定义的 stringer
func (v *Duck) String() string {
	return fmt.Sprintf("\nMy name is %v", v.name)
}

// 在 Go 内，error 也是一个常用的 interface
func (v *Duck) Error() string {
	return fmt.Sprintf("%v caught error!", v.name)
}

// 与其他语言不同，Go 允许你在函数内判断 nil 并处理，而不是直接抛出 null pointer exception
func (v *Duck) whisper(data string) {
	if v == nil {
		fmt.Printf("<nil>")
		return
	}
	fmt.Println(v.name, "is saying:", data)
}

func (v *Duck) getName() string {
	return v.name
}

func describe(data Talk) {
	fmt.Printf("value is %v, and type is %T\n", data, data)
}

// 用 interface 创建一个空对象（类似 javascript `{}`）
func declareEmpty() {
	v := map[string]interface{}{}
	fmt.Printf("\ntypeof v is %T, and value is %v", v, v)
}

func main() {
	//quickDuckInterface()
	//testNil()
	//emptyInterface()
	switchInterface("12")
	v := Duck{"Yellow duck"}
	fmt.Println(v.String())
	errorDuck := &Duck{"ErrorDuck"}
	if errorDuck != nil {
		fmt.Printf(errorDuck.Error())
	}
	declareEmpty()
}
