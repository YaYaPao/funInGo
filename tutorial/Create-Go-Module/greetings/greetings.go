package greetings

import "fmt"

// Hello 在 go 中，如果一个 package 内的函数名首字母大写，则表示该函数能够被其他 package 调用（类似 export ）
// 在 go 内 := 表示初始化变量且为其赋值
func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
