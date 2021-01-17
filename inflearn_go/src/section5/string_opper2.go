//데이터 타입 문자열 연산

package main

import "fmt"

func main() {
	//추출

	var str1 string = "Golang"
	var str2 string = "World"

	fmt.Println(str1 == str2)
	fmt.Println(str1 != str2)
	fmt.Println(str1 > str2) //go 에선 문자열 비교할 때 아스키 코드에 대한 사전식 비교를 한다.

}
