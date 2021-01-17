//자료형 슬라이스 심화
package main

import "fmt"

func main(){

	//슬라이스 복사
	//copy
	//make로 반드시 용량을 할당 후 복사해야 한다.
	//복사된 슬라이스 값 변경해도 원본에는 영향 없음

	//예제1
	slice1 := []int{1,2,3,4,5,6,7,8,9,10}
	slice2 := make([]int, 5)
	slice3 := []int{}

	copy(slice2, slice1)
	copy(slice3, slice1)

	fmt.Println(slice2)
	fmt.Println(slice3) //복사가 안됨

	//예제2
	a := []int{1,2,3,4,5}
	b :=make([]int, 5)

	copy(b, a)
	b[0] = 7

	fmt.Println(a)
	fmt.Println(b) //복사가 안됨

	//예제 3
	c := [5]int{1,2,3,4,5}
	d := c[0:3] //부분적 슬라이스 추출은 참조이다. 원본값이 변경 된다.

	d[1] = 7

	b[0] = 7

	fmt.Println(c)
	fmt.Println(d) 

	//예제 4
	e := []int{1,2,3,4,5,6,7,8,9}
	f := e[0:5:7] //세번째 값은 용량 선언

	fmt.Println(len(f), cap(f), f)




}