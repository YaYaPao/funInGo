package main

import "fmt"

// channels 通过管道符 `<-` 来传递数据，通过配置 goroutine 使用
// 必须通过 `make(chan int)` 来创建
// 默认会等到对端的任务执行完毕并返回结果后再进行数据通信，结合 javascript 的 async/await 就能理解

// 通过 `c <- sum` 将 sum 值传递给管道
func sum(data []int, c chan int) {
	sum := 0
	for _, v := range data {
		sum += v
	}
	fmt.Println(sum)
	c <- sum
}

// 如果缓冲区满了，再往里面写入会抛出异常 `fatal error: all goroutines are asleep - deadlock!`
func channelBuffer() {
	c := make(chan int, 3)
	c <- 1
	c <- 2
	c <- 3
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}

// 通过 `go` 创建了两个 goroutine，并将管道传递给 sum 函数，在其处理完毕之后，通过管道将值发送
// 通过 `<-c` 取回管道内传输值
func channelTour() {
	c := make(chan int)
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5, 6}
	go sum(s1, c)
	go sum(s2, c)
	x, y := <-c, <-c
	fmt.Printf("x is %v, y is %v", x, y)
}

// 如果确认不会有新的值需要发送后，可以在发送端主动关闭 channel
// channel 不像文件，一般不建议主动关闭 channel
// ⚠️不允许在接受端关闭，向一个关闭的 channel 发送数据会报错，这个动作只能由发送端执行
func fibonacci(data int, c chan int) {
	x, y := 0, 1
	for i := 0; i < data; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

// 因为 channel 一直在返回数据，因此可以对其进行遍历，直到当前 channel 被关闭
func closeChannel() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

func selectChannel(c chan int, quit chan int) {
	x := 0
	for {
		select {
		case c <- x:
			x++
		case <-quit:
			fmt.Println("quit")
			return
		default:
			fmt.Println("ddd")
		}
	}
}

// 通过 select 在 channel 内匹配不同的 case 进行执行，不同于 switch，这里 case 可以看作是一个传输关系
//  - select 在匹配到 case 之前会挂起
//  - 如果多个 case 同时匹配，会随机挑选一个执行
func selectChannelGuide() {
	c := make(chan int, 100)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	selectChannel(c, quit)
}

func main() {
	//channelTour()
	//channelBuffer()
	//closeChannel()
	selectChannelGuide()
}
