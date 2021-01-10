package main

import "fmt"

func main()  {
	//제어문(조건문)
	//if문은 반드시 boolean 형으로 검사 할 것 -> 1, 0(사용불가 : 자동형 변환 불가)
	//if문은 소괄호 사용 x
	//if문 (1)
	var a int = 20
	b := 30

	if a >= 15{
		fmt.Println("15이상")
	}

	if b>= 25{
		fmt.Println("25이상")
	}

	if c := 40; c >=35{
		fmt.Println("35명 이상")
	}
}