// 함수 기초(1)
package main

import "fmt"
//함수 선언 위치는 어느 곳이든 가능




func muti_fly(n ...int) int{
	tot := 1
	for _, value := range n{
		tot *= value
	}

	// fmt.Println(tot)
	return tot
}

func sum(n ...string) {
	sum := n
	for _, value := range sum{
		fmt.Println(value)

	}

}
func main(){
	//함수 고급
	//가변 인자 실습(매개 변수 개수가 동적으로 변할때) -> 정해지지 않을때
	x:= muti_fly(5, 6, 7, 8)

	fmt.Println(x)

	sum("a", "apple")


	a := []int{5,6,7,8,9,10}

	m:= muti_fly(a...) //슬라이스를 넣을 때 이렇게 넣으면 편함...
	fmt.Println(m)

}