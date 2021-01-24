// 함수 기초(1)
package main

import "fmt"
//함수 선언 위치는 어느 곳이든 가능




func fact(n int) int {
	if n ==0{
		return 1
	}
	return n*fact(n-1)
}

func main(){
	//재귀 함수
	//재귀 함수(recursion)
	//프로그램이 보기 쉽고, 코드 간결, 오류 수정이 용이
	//코드를 이해하기 어려움

	//예제 1
	x:=fact(7)

	fmt.Println(x)

}