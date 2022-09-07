package greetings

import (
	"errors"
	"fmt"
)

// Hello 在 go 中，如果一个 package 内的函数名首字母大写，则表示该函数能够被其他 package 调用（类似 export ）
// 在 go 内 := 表示初始化变量且为其赋值
// 通过返回 error|nil 来了解 go 中如何处理错误 -> "Handling errors is an essential feature of solid code"
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Empty name")
	}
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
