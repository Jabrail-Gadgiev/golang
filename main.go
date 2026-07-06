package main

import (
	"fmt"
)

func main() {
	// [inclusive:exclusive]
	xi := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x1 := xi[:5]
	x2 := xi[5:]
	x3 := xi[2:7]
	x4 := xi[1:6]
	fmt.Printf("%v\n%v\n%v\n%v\n", x1, x2, x3, x4)
}
