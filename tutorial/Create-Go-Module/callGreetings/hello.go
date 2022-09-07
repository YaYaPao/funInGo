package main

// 引入本地相对路径的 module，我们需要：
// 1. `go mod edit --replace greetings=../greetings`，用来指定 greetings 应该由 ../greetings 来定位；
// 2. `go mod tidy` 同步 module 依赖
// 在生产环境，我们需要将 module 发布后进行拉取
import (
	"fmt"
	"greetings"
	"log"
)

func main() {
	// 设置 logger 的前缀以及关闭其对时间、源文件和行数的打印
	log.SetPrefix("greetings:")
	log.SetFlags(0)

	// call Hello method
	//message, err := greetings.Hello("papaya")

	// call Hellos method
	names := []string{"glory", "young", "doggy"}
	message, err := greetings.Hellos(names)

	// 如果返回错误，则将错误打印在控制台，并且退出程序
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
