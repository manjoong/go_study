package main

import (
	"fmt"
)

func main() {

	var n1 uint8 = 125
	var n2 uint8 = 90
	// var n3 uint32 = math.MaxUint32
	// var n4 uint64 = math.MaxUint64

	fmt.Println(n1 + n2)
	fmt.Println(n1 - n2)
	fmt.Println(n1 * n2)
	fmt.Println(^n1)

}
