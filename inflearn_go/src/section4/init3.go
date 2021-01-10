package main

import (
	"fmt"
)

var num int32

func main() {
	//init : 패키지 로드시에 가장 먼저 호출
	//가장 먼저 초기화 되는 작업 작성 시 유용하다.

	fmt.Println("Main Method Start!")
}

func init() {
	//init : 패키지 로드시에 가장 먼저 호출
	//가장 먼저 초기화 되는 작업 작성 시 유용하다.
	num = 30
	fmt.Println("Init Method Start!")
}
