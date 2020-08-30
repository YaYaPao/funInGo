package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func testMath() {
	var n1 float64 = 7
	var n2 float64 = 6
	var n3 float64 = 3.74
	fmt.Printf("max value is %v\n", math.Max(n1, n2))
	fmt.Printf("pow value is %v\n", math.Pow(n2, 2))
	// Trunc 取整
	fmt.Printf("int value is %v\n", math.Trunc(n3))
}

// 随机数生成
func testRand() {
	fmt.Println("-------- rand --------")
	seed := time.Now().UnixNano()
	fmt.Println(seed)
	// 生产一个随机数，默认 seed 为 1，因为 seed 没有发生改变，所以 r1 实际上是随机生成的一个不变的数字
	r1 := rand.Int()
	fmt.Println(r1)
	// 指定 seed，当前 seed 随着时间变化，因此可以观察到，不同于 r1，r2 每次都会发生变化，即动态随机数
	rand.Seed(seed)
	r2 := rand.Int()
	fmt.Println(r2)
	// 返回 0-10 内的整型随机数
	r3 := rand.Intn(10)
	fmt.Println(r3)
	// 返回 m-(m+10) 范围内的随机数
	r4 := rand.Intn(10) + 5
	fmt.Println(r4)
}

func main() {
	testMath()
	testRand()
}
