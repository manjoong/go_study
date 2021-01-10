//데이터 타입 : bool(1)
package main

import "fmt"

func main() {
	//boolean -> 참과 거짓
	//조건부 논리 연산자랑 주로 사용 | ||(or) &&(and)
	//암묵적 형 변환x 0n nil != false

	var b1 bool = true
	var b2 bool = false
	b3 := true

	fmt.Println("ex1: ", b1)
	fmt.Println("ex1: ", b2)
	fmt.Println("ex1: ", b3)

	if b3 {
		fmt.Println("b3 is trure")

	}
}
