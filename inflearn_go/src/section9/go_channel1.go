package main

import (
	"fmt"
	"time"
)

func work1 (v chan int) {
	fmt.Println("work1 : s --->", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("work1 : e --->", time.Now())
	v <- 1

}

func work2 (v chan int) {
	fmt.Println("work2 : s --->", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("work2 : e --->", time.Now())
	v <- 2

}

func main(){
	//고루틴간의 상호 정보(데이터) 교환 및 실행 흐름 동기화 위해 사용

	//실행 흐름 제어가능 -> 일반 변수로 선언 후 사용가능
	//데이터 전달 자료형 선언 후 사용(지정된 타입만 주고 받을 수 있음)
	//interface{}는 모든 자료형을 받을 수 있음 
	//값을 전달 복수 후, bool. int, 포인터(슬라이스, 맵) 전달시에는 주의. 동기화 사용
	//멀티 프로세싱 처리에서 교착상태(경쟁) 주의!
	// <-, -> (채널 <- 데이터) => 데이터를 채널로 보내겠다 :송신
	// 채널 -> 데이터 :수신

	//예제1

	fmt.Println("Main: s --> ", time.Now())

	// var c chan int
	// c = make(chan int)

	v := make(chan int) //int형 채널 선언

	go work1(v)
	go work2(v)

	<- v
	<- v
	fmt.Println("Main: e --> ", time.Now())
}