package main

import "fmt"

func main() {

	//부동 소수점

	var num1 float32 = 0.14
	var num2 float32 = .7563
	var num3 float32 = 442.4244
	var num4 float32 = 10.0

	var num5 float32 = 14e6
	var num6 float64 = .15232

	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(num3)

	fmt.Println(float32(num4 - 0.1))
	fmt.Println(float64(num4 - 0.1))
	fmt.Println(num5)
	fmt.Println(num6)
}
