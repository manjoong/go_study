package main

import "fmt"


func ex_f1(){
	fmt.Println("fi: start!")
	defer ex_f2()
	fmt.Println("f1: end!")
}

func ex_f2(){
	fmt.Println("f2: called")
}
func main()  {
	//defer 지연실행 (미뤄놨다가 나중에 실행)
	//defer를 호출한 함수가 종료되기 직전에 호출 된다
	// 타 언어의 finally문과 비슷
	//주로 리소스 반한, 열린 파일
	ex_f1()
	
}