package main

import "fmt"

//슬라이스는 nil값이 있다. 이 nil값은 마치 비어 있는 것처럼 처리 되나 nil을 출력한다고 해서 오류가 나진 않는다
func main() {
	var intSlice []int
	var stringSlice []string
	fmt.Println(intSlice)
	fmt.Println(stringSlice)
	fmt.Println(len(intSlice))
	slice := append(intSlice, 1)
	fmt.Println(slice)
}