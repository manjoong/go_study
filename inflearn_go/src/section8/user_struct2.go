package main

import "fmt"

type cnt int

func main()  {
	//기본 자료형 사용자 정의 타입
	//예제1

	a:= cnt(15)
	fmt.Println(a)
	//예제2

	var b cnt = 15
	fmt.Println(b)

	//testConvT(b) 예외 발생 사용자 정의 타입<-> 기본타입: 매개변수 전달 시에 변환해야 사용 가능
	testConvT(int(b))
	testConvD(b)


	//예제3

	
}

//cnt는 결국 Int니깐 둘다 실행 가능?
func testConvT(i int){
	fmt.Println(i)
}

func testConvD(i cnt){
	fmt.Println(i)
}