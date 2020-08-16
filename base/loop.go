package main

import "fmt"

func loopOne() {
	for i := 0; i < 7; i++ {
		fmt.Println("i is", i)
	}
	// 省略初试语句、判断条件的写法
	j := 0
	for ; ; j++ {
		if j > 3 {
			// 通过 break 跳出循环
			break
		}
		fmt.Println("j is", j)
	}
}

func loopTwo() {
	var a int
	// 类 while 写法
	for a < 3 {
		fmt.Println("a is", a)
		a++
	}
}

func loopThree() {
	var b int = 7
	// 通过 break 跳出循环，结束语句放在循环体内
	for {
		if b < 3 {
			break
		}
		fmt.Println("b is", b)
		b--
	}
}

func loopRange() {
	s := "hello"
	for i, value := range s {
		fmt.Printf("index is %d, ASCII value is %d, chart is %c \n", i, value, value)
	}
}

func traggle() {
	lines := 6
	for i := 0; i < lines; i++ {
		for n := 0; n < i+1; n++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}

func main() {
	loopOne()
	loopTwo()
	loopThree()
	loopRange()
	traggle()
}
