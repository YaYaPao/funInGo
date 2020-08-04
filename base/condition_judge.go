package main

import "fmt"

func test() {
	source := 20
	if source % 2 == 0 {
		fmt.Printf("%v is 偶数 \n", source)
	} else {
		fmt.Printf("%v is 奇数 \n", source)
	}

	// 在条件语句内声明变量
	if a := 21; a % 2 == 1 {
		fmt.Printf("%v is male \n", a)
	} else {
		fmt.Printf("%v is female \n", a)
	}
}

func testSwitch() {
	source := 20
	grade := ""
	// 默认为 true，通过条件语句进行判断
	switch {
	case source > 10:
		grade = "A"
	case source > 20:
		grade = "B"
	}
	fmt.Printf("grade is %v", grade)

	//
}

func main() {
	test()
	testSwitch()
}
