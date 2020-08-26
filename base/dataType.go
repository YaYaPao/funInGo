package main

import "fmt"

var temp string = `
hello, world!
i am an engineer!
multi strings
`

func testString() {
	a := "hello你好"
	for i := 0; i < len(a); i++ {
		// %x 十六进制，小写字母，每字节两个字符
		fmt.Printf("%x", a[i])
		// %c 相应的 Unicode 码点所表示的字符
		fmt.Printf("%c", a[i])
		// 字节类型的值
		fmt.Println(a[i])
	}

	// 解决乱码问题，涉及到中文，则以 rune 格式进行解析
	r := []rune(a)
	fmt.Println(r)
	for i := 0; i < len(r); i++ {
		fmt.Printf("%x", r[i])
		fmt.Printf("%c", r[i])
		fmt.Println(r[i])
	}
}

func main() {
	fmt.Println(temp)
	var a byte = 'a'
	var b rune = '一'
	// 输出值对应的 unicode
	fmt.Println(a, b)
	testString()
}
