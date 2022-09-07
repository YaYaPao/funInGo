package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello 在 go 中，如果一个 package 内的函数名首字母大写，则表示该函数能够被其他 package 调用（类似 export ）
// module 内方法名如果首字母小写，则不会暴露出去
// 在 go 内 := 表示初始化变量且为其赋值
// 通过返回 error|nil 来了解 go 中如何处理错误 -> "Handling errors is an essential feature of solid code"
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Empty name")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// go 会在程序初始化时自动执行 `init()` 方法
// 设置 rand 对象的初始值以保证 rand 方法的随机性
func init() {
	fmt.Println("init start")
	rand.Seed(time.Now().UnixNano())
}

// 学习定义 slice 的方式以及随机值的使用
// 这里设置 slice 为 []string 来告诉 go，该变量的底层 size 可能会动态变化
func randomFormat() string {
	greetingsMessage := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v.",
		"Hail, %v. Well met!",
	}
	return greetingsMessage[rand.Intn(len(greetingsMessage))]
}
