package main

import (
	"fmt"
	"strings"
)

// io 包定义了 `io.Reader` 接口，其中 `Read` 在 Go 内经常用到，用来处理流式数据传输
// `Read` 通过填充一个定义好的 byte splice，同时返回 error(如果无数据进行填充，会返回 `io.EOF` 错误)

// 通过 `strings.NewReader()` 创建一个新的流
// 定义一个 byte slice 作为容器
// 通过 `r.Read()` 来分批读取流式数据
// 当捕获到 EOF 错误时，退出循环
func main() {
	r := strings.NewReader("hello reader!")
	var s = make([]byte, 6)
	for {
		n, err := r.Read(s)
		fmt.Printf("Read size is %v, error is %v\n", n, err)
		fmt.Printf("current read content is %q\n", s[:n])
		if err != nil {
			fmt.Println("Caught error: ", err)
			break
		}
	}
}
