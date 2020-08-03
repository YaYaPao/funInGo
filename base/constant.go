package main
import "fmt"

func main() {
	const W, H = 10, 20
	fmt.Println(W, H)
	// 模拟枚举
	const (
		MALE = 0
		FEMALE = 1
	)
	fmt.Println(MALE, FEMALE)
	// b,c 的值被自动设为第一个有值的常量值
	const (
		a = 10
		b
		c
	)
	fmt.Println(a, b, c)
	const (
		d = iota
		e
	)
	// 在新的 const() 内 iota 重置为 0
	const (
		f = iota
		g
	)
	fmt.Println(d, e, f, g)
}
