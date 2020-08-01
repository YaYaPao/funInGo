package main

import "fmt"

func testVariable () {
	var a int = 7
	var b = "hello"
	c := true
	var (
		d int
		e string
		f func()
	)
	fmt.Print(a, b, c, d, e, f)
}

func main() {
	testVariable()
}
