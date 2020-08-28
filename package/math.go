package main

import (
	"fmt"
	"math"
)

func testMath() {
	var n1 float64 = 7
	var n2 float64 = 6
	fmt.Printf("max value is %v\n", math.Max(n1, n2))
	fmt.Printf("pow value is %v\n", math.Pow(n2, 2))
}

func main() {
	testMath()
}
