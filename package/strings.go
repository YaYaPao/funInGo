package main

import (
	"fmt"
	"strings"
	"unicode"
)

func searchString() {
	fmt.Println("------Search String-------")
	a := "hello"
	b := "world"
	// 从词典上比较两个字符串
	com := strings.Compare(a, b)
	fmt.Println("compare result is ---", com)
	hasValue := strings.Contains(a, "e")
	fmt.Printf("%v has e is %v\n", a, hasValue)
	count := strings.Count(a, "l")
	fmt.Printf("%v has %v l\n", a, count)
}

func splitString() {
	fmt.Println("------Split String-------")
	s := "hello world!"
	// 切割后返回一个 slice
	arr := strings.Split(s, "o")
	fmt.Println(arr)
	for _, v := range arr {
		fmt.Printf("value is %v and it's length is %d\n", v, len(v))
	}
	arr2 := strings.SplitAfter(s, "o")
	fmt.Println(arr2)
	arr3 := strings.FieldsFunc(s, func(r rune) bool {
		return string(r) == "l"
	})
	fmt.Println(arr3)
}

func transString() {
	fmt.Println("------Trans String-------")
	s := "hello World!"
	fmt.Printf("首字母大写：%v\n", strings.Title(s))
	fmt.Printf("全小写：%v\n", strings.ToLower(s))
	fmt.Printf("全大写：%v\n", strings.ToUpper(s))
	fmt.Printf("指定字符替换：%v\n", strings.ToLowerSpecial(unicode.TurkishCase, s))
}

func trimString() {
	fmt.Println("------Trim String-------")
	s := "  hello world@@@ "
	fmt.Printf("去掉首尾空格：%v\n", strings.TrimSpace(s))
	s1 := "hello@@@@"
	fmt.Printf("去掉尾部指定字符：%v\n", strings.TrimRight(s1, "@"))
}

func compareString() {
	fmt.Println("------Compare String-------")
	// JavaScript 中常用的 join
	arr := []string{"h", "e", "l", "l", "o"}
	s := strings.Join(arr, ",")
	fmt.Printf("typeof %v is %T\n", s, s)

	s1 := "world"
	// 按照字典顺序比较两个字符串的大小，返回 -1， 0， 1
	fmt.Printf("compare result is %v\n", strings.Compare(s, s1))
	// 比较两个字符串是否相等，忽略大小写，为否则返回 -1
	fmt.Printf("equal>>>> %v\n", strings.EqualFold(s, s1))
	// 返回字符串重复 n 次返回
	fmt.Printf("repeat times is %v\n", strings.Repeat(s, 3))
	// 替换指定字符串, n 小于 0 表示全部替换
	fmt.Printf("replace l with q: %v\n", strings.Replace(s, "l", "q", -1))
}

func main() {
	// test strings methods
	// you can link https://golang.org/pkg/strings/#Compare for more details
	// just have fun!
	searchString()
	splitString()
	transString()
	trimString()
	compareString()
}
