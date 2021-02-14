//인터페이스 기본

package main


import "fmt"

type Dog struct{
	name string
	weight int
}

type Cat struct{
	name string
	weight int
}

func printValue(s interface{}){
	fmt.Println(s)
}

func main()  {
	// 인터페이스 활용(빈 인터페이스)
	//함수 내에서 어떠한 타입이라도 유연하게 매개변수로 받을 수 있다 (만능) -> 모든 타입 지정 가능

	dog:=Dog{"poll", 10}
	cat:=Cat{"nami", 5}

	//인터페이스는 모든 자료형을 다 받을 수 있음
	printValue(dog)
	printValue(cat)
	// act(dog1)
	// act(cat1)

	printValue(15)
	printValue("Animall")
	printValue([5]Dog{})
	printValue(false)

	

}