package main

import "fmt"

// `main` 是当前 package 的入口，当执行 `go run` 时会被默认执行
// 执行 `go run` 并不会生成一个可执行的二进制文件，而是会直接编译并执行
// 在 rootDir 下，执行 `go run tutorial/`
// 通过 `go mod init tutorial/hello.go` 来管理 package 依赖

func main() {
	fmt.Println("Hello, world!")
}
