package main

import (
	"fmt"
	"strconv"
	"strings"
)

func IsNumString(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func main() {
	// 字符串转小写
	str := strings.ToLower("HELLO")
	fmt.Println(str)

	is_num := IsNumString("123a")
	fmt.Println(is_num)
}
