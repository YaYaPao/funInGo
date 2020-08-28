package main

import (
	"fmt"
	"time"
)

func timeDisplay() {
	fmt.Println("------ time display ------")
	// 当前本地时间
	t := time.Now()
	fmt.Println(t)
	// 返回时间实例的时间戳，seconds
	u := t.Unix()
	fmt.Printf("%v timestamp is %v\n", t, u)
	// 返回指定时间的年，月，日
	y, m, d := t.Date()
	fmt.Printf("yaer is %v and month is %v and day is %v\n", y, m, d)
	// time 格式转换成 string 类型
	s := t.String()
	fmt.Printf("typeof t is %T and typeof s is %T\n", t, s)
	// format time, more attention, the layout should be 2006/01/02 15:04:05
	// 据说是因为此时为 go 的诞生日，任性！
	f := t.Format("2006/01/02 15:04:05")
	fmt.Printf("the value is %v and type is %T\n", f, f)
	// parse 用来将字符串转为 time.Time 格式，layout 用来对齐字符串格式
	f2, err := time.Parse("2006/01/02 15:04:05", f)
	if err == nil {
		fmt.Println(f2)
	}
}

func main() {
	timeDisplay()
}
