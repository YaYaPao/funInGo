package main

// 引入本地相对路径的 module，我们需要：
// 1. `go mod edit --replace greetings=../greetings`，用来指定 greetings 应该由 ../greetings 来定位；
// 2. `go mod tidy` 同步 module 依赖
// 在生产环境，我们需要将 module 发布后进行拉取
import (
	"fmt"
	"greetings"
)

func main() {
	message := greetings.Hello("papaya")
	fmt.Println(message)
}
