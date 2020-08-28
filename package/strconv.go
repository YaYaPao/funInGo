package main

import (
	"fmt"
	"strconv"
)

func parseString() {
	s := "5267"
	// attention Atoi will return multiple-value, so you should multiple-variable to payload
	// conversion from string to int
	// 如果不是一个整型字符串，则会返回 0
	i, _ := strconv.Atoi(s)
	fmt.Printf("test Atoi and the result is %v and it's type is %T\n", i, i)

	// 指定当前整型字符串的进制，同时指定返回的 int 类型，注意必须是无溢出的整数类型
	num1, _ := strconv.ParseInt("0101", 2, 64)
	fmt.Printf("the 0101 conversions from 2 to 10 result is: %v\n", num1)
	// 这里将 0101 当成 10 进制进行解析
	num2, _ := strconv.ParseInt("0101", 10, 64)
	fmt.Printf("the 0101 conversions from 10 to 10 result is: %v\n", num2)

	// 返回字符串表示的布尔值，它接受1, t, T, TRUE, true, 0, f, false, FALSE，其他字符串会返回错误
	b1, _ := strconv.ParseBool("t")
	fmt.Printf("result is %v\n", b1)
	// test error
	_, e1 := strconv.ParseBool("hello")
	fmt.Printf("error message is ------ %v\n", e1)
}

// Format 主要是将其他类型数据转换为字符串类型
func formatString() {
	fmt.Println("------ format string ------")
	// int to string
	i := 567
	s1 := strconv.Itoa(i)
	fmt.Printf("coversions from int to string: type is %T and value is %v\n", s1, s1)

	// bool to string: true or false
	b := true
	s2 := strconv.FormatBool(b)
	fmt.Printf("conversions from bool to string: type is %T and value is %v\n", s2, s2)
}

func main() {
	parseString()
	formatString()
}
