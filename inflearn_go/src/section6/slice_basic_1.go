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

	//길이가 가변 -> 동적으로 크기가 늘어난다. 레퍼런스(참조 값) 타입
	//슬라이스(길이 & 용량) 크기가 동적으로 할당 가능

	//2가지 선언 방법 1. 배열처럼 선언 2. make(자료형, 길이 용량)함수 선언

	var slice1 = []int{}
	slice2 := []int{}
	slice3 := []int{1,2,3,4,5}
	slice4 := [][]int{
		{1,2,3,4,5},
		{6,7,8,9,10},
	}
	// slice2[0] = 1
	slice3[4]=10

	fmt.Printf("%-5T %d %d %v\n", slice1, len(slice1), cap(slice1), slice1)
	fmt.Printf("%-5T %d %d %v\n", slice2, len(slice2), cap(slice2), slice2)
	fmt.Printf("%-5T %d %d %v\n", slice3, len(slice3), cap(slice3), slice3)
	fmt.Printf("%-5T %d %d %v\n", slice4, len(slice4), cap(slice4), slice4)

	//예제 2
	var slice5 []int = make([]int, 5, 10) //용량에 근접하게 잡을수 있는게 가장 예측하여 사용할 수 있기 때문에 size를 특정한다.그냥 사이즈를 빼면 용량도 5로 선언
	var slice6 = make([]int, 5)
	slice7 :=make([]int, 5, 100)
	slice8 :=make([]int, 5)

	slice6[2] = 7 //make로 만들면 0으로 초기화 해주기 때문에 해당 명령어가 먹음

	fmt.Printf("%-5T %d %d %v\n", slice5, len(slice5), cap(slice5), slice5)
	fmt.Printf("%-5T %d %d %v\n", slice6, len(slice6), cap(slice6), slice6)
	fmt.Printf("%-5T %d %d %v\n", slice7, len(slice7), cap(slice7), slice7)
	fmt.Printf("%-5T %d %d %v\n", slice8, len(slice8), cap(slice8), slice8)
	//그냥 배열처럼 선언할때랑 make로 처리할때 다른점 => make는 0으로 초기화 해준다.

	//예제3
	var slice9 []int //nil 슬라이스(길이와 용량이 0)
	if slice9 == nil{
		fmt.Println("slice9은 nil")
	}

}