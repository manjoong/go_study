//구조체 기본

package main

import "fmt"

type Account struct {
	number string
	balance float64
	interest float64
}
//리시버를 통해서 구조체<->메서드를 연결

func (a Account) Calculate() float64{
	return a.balance + (a.balance* a.interest)
}

func  main()  {
	//구조체
	//go의 필드들의 집합체 또는 컨테이너
	//필드는 갖지만 메소드는 갖지 않는다.
	//객체지향 방식을 지원 -> 리시버를 통해 메소드랑 연결
	//상속, 객체, 클래스 개념 없음
	//구조체 -> 구조체 선언 -> 구조체 인스턴스(리시버)

	//예제1
	kim := Account{number: "245-601", balance: 10000000, interest: 0.015}
	lee := Account{number: "245-421", balance: 12000000}
	park := Account{number: "245-903", interest: 0.025}
	cho := Account{"245-903", 150000, 0.03}

	fmt.Println(kim)
	fmt.Println(lee)
	fmt.Println(park)
	fmt.Println(cho)

	//예제2

	fmt.Println(int(kim.Calculate()))

	
}