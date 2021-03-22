package main

import "fmt"

func runFunc(){
	defer func(){
		s := recover()
		fmt.Println("Error message", s)
	}()

	panic("Error occured!")
	fmt.Println("test")
}
func main(){
	//go panic 함수
	//에러 복구 가능
	//panic으로 발생한 에러를 복구 후 코드 재 실행
	//즉 코드 흐름을 정상상태로 복구 하는 기능

	runFunc()

	fmt.Println("end main")

}