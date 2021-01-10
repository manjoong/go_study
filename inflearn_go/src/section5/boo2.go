package main

import "fmt"

func main() {
	//boolean -> 참과 거짓
	//조건부 논리 연산자랑 주로 사용 | ||(or) &&(and)
	//암묵적 형 변환x 0n nil != false

	fmt.Println("ex1: ", true && true)
	fmt.Println("ex1: ", true && false)
	fmt.Println("ex1: ", false && false)

	num1 := 15
	num2 := 37

	fmt.Println("ex1: ", num1 < num2)

}
