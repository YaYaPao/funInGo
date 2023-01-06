package main

import "fmt"

// 在 go 内，`struct` 是字段的集合

type Heros struct {
	A string
	B string
}

// 结构体继承
type SuperMan struct {
	Heros
	Name string
}

// 通过 `struct{}` 的形式来初始化，通过 `.` 来进行取值（或者重新赋值）
// 在 go 内通过 `&` 对 structs 进行取址后，不再需要用 `*` 来进行取值操作，go 会默认处理
// 在 Go 内，对于 struct, map 都应使用取址符，避免 cant assigned error
func main() {
	h := Heros{"PA", "UG"}
	h.A = "SA"
	p := &h
	p.B = "JUGG"
	fmt.Printf("Heros are %v\n", h)
	fmt.Println(p)
}
