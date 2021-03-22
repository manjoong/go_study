//go 에러 처리 기초(2)

package main

import (
	"fmt"
	"log"
)

func notZero(n int)(string, error){ //메소드 리턴 값 error타입 중요!
	if n !=0{
		s:=fmt.Sprint(n)
		return s, nil
	}

	return "", fmt.Errorf("%d를 입력했습니다. 에러 발생!", n)
}

func main(){
	//에러 처리
	//Errorf 를 이용한 에러 처리

	a, err := notZero(1)

	//기본적인 메서드 에러 처리 예제
	if err != nil {
		// log.Fatal(err.Error())//방법 1
		log.Fatal(err)//방법 2
	}

	fmt.Println(a)

	b, err := notZero(0)
	if err != nil {
		// log.Fatal(err.Error())//방법 1
		log.Fatal(err)//방법 2
	}

	//fatal 이후 프로그램 종료 후 실행 중지

	fmt.Println(b)
	fmt.Println("end error handling")
}