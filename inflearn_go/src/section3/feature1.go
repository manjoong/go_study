package main

import "fmt"

func main(){
	//go는 모호하거나, 애매한 문법 금지
	//후치 연산자 허용 i++, 전위(전치) 연산자는 비허용
	//문법 에러, 증감 연산-> 반환값 없음
	//포인터(pointer)를 허용, 연산 비허용

	sum, i := 0, 0
	for i <=100{
		//sum += i++ 예외 발생
		sum += i
		i++
		// ++i 전위증가문 예외발생
	}
	fmt.Println(sum)

}