//데이터 타입 문자열 연산

package main

import "fmt"

func main() {
	//추출

	var str1 string = "Golang"
	var str2 string = "World"

	fmt.Println(str1[0:2], str1[0])
	fmt.Println(str2[3:], str2[0])
	fmt.Println(str2[:4])

}
