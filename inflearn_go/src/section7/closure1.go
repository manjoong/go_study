package main

import "fmt"

func main()  {
	//클로저
	//익명함수 사용할 경우 함수를 변수에 할당해서 사용 가능
	//함수 안에서 함수를 선언 및 정의 가능 -> 이 때 외부 함수에 선언 된 변숭 접근 가능한 함수
	//함수가 선언되는 순간에 함수가 실행 될 때 실체의 외부 변수에 접근하기 위한 스냅샷(객체)

	//함수를 호출 할 때 이전에 존재했던 값을 유지하기 위해서 ㅠ
	multiply := func(x, y int) int{
		return x*y
	}

	r1:= multiply(5, 10)

	fmt.Println(r1)

	//예제2

	m, n := 5, 10 //변수가 캡쳐
	sum := func(c int) int{ //익명함수가 변수 할당
		return m+n+c //지역 변수가 소멸 되지 않는다.
	}

	r2 := sum(10)
	fmt.Println(r2)
	
}