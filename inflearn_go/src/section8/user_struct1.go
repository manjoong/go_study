package main

import "fmt"

type Car struct {
	name string
	color string
	price int
	tax int
}

//구조체 <-> 메소드 바인딩

//일반 메서드
func Price(c Car) int{
	return c.price + c.tax
}

//구조체 <-> 메소드 바인딩
func (c Car) Price() int{
	return c.price+c.tax
}

func main(){
	//go 에서는 객체지향타입을 구조체로 정의 한다(클래스 상속 개념 없음)
	// 객체 지향-> 클래스(속성, 기능(상태 : 메소드)) : 코드의 재사용성, 코드의 관리가 용이, 신회성이 높은 프로그래밍
	//객체지햐 언어일까?
	//go는 젆ㅇ적인 객체지향의 특징을 가지고 있지 않지만, 인터페이스 -> 다형성 지원, 구조체를 클래스 형티의 코딩 가능
	// 갹채지향의 기본 개념 -> go 에서 포함하고 있다 -> 객체 지향 프로그래밍 언어
	// 상태, 메소드를 분리해서 정의(결합성 없음)
	//사용자 정의 타입: 구조체, 인터페이스, 기본 타입
	//구조체와 메소드 연결을 통해서 타 언어의 클래스 형식처럼 사용 가능(객체 지향)

	//예제 1

	bmw := Car{name: "520d", price: 5000000, color: "white", tax: 5000}
	benz := Car{name: "220d", price: 6000000, color: "white", tax: 6000}

	fmt.Println(bmw, &bmw)
	fmt.Println(benz, &benz)

	fmt.Println(Price(bmw)) //-> 보통 이렇게 선언하지 않음. 이건 구조체에 메서드가 속한게 아니라 그냥 변수를 대입한 함수
	fmt.Println(bmw.price)



	
}