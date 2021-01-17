//자료형 배열

package main

import "fmt"

func main(){
	//배열
	//배열은 용량과 길이가 같다.
	//배열 vs 슬라이스
	//길이가 고정 vs 길이가 가변
	//값 타입 vs 참조 타입
	//복사 전달 vs 참조 값 전달
	//전체 비교 연산자 사용 가능 vs 비교 연산자 사용 불가능
	//대부분 슬라이스를 사용한다.
	//cap() 배열, 슬라이스의 용량
	//len() 배열 슬라이스 개수


	//예제1
	var arr1 [5]int
	var arr2 [5]int = [5]int{1,2,3,4,5}
	arr3 := [5]int{1,2,3,4,5}
	arr4 := [5]int{1,2,3}
	arr5 := [...]int{1,2,3} //배열 크기 자동 맞춤

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr4)
	fmt.Println(arr5)

	arr1[2] = 5
	fmt.Printf("%d %v\n", len(arr1), arr1)
	

}