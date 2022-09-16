package main

import (
	"fmt"
	"time"
)

// Goroutine 是 Go 内一个轻量的线程管理工具，通过 `go` 关键字进行启动一个新的线程执行指定函数
// Goroutine 运行在相同的地址空间上，因此必须同步访问共享内存

func sayHello(data string) {
	fmt.Println(time.Millisecond)
	for i := 0; i < 3; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(data)
	}
}

// `go sayHello` 会启动一个新的线程来执行 `sayHello()`
//
func main() {
	go sayHello("i am Mike.")
	sayHello("i am Jack")
}
