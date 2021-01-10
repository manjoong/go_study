package main

import "fmt"

func main(){
	//열거형
	//상수를 사용하는 일정한 규칙에 따라 숫자를 계산 및 증가시키는 묶음
	//쓸만하네..
	const(
		_ = iota
		DEFAULT
		SILVER
		GOLD
	)

	fmt.Println(DEFAULT)
	fmt.Println(SILVER)
	fmt.Println(GOLD)
}