package main

import "fmt"

func main(){
	//go panic 함수
	//사용자가 에러 발생 가능
	//panic 함수는 호출 즉시, 해당 메서드를 즉시 중지 시키고 defer함수를 자기자신을 호출한 곳으로 리턴함

	//런타임 이외에 사용자가 코드 흐름에 따라 에러를 발생 시킬때 중요.
	//문법적인 에러는 아니지만, 논리적인 코드 흐름에 따른 에러 발생 처리 가능

	fmt.Println("Start Main")
	panic("Error occured: user panic!")

	fmt.Println("end main")

}