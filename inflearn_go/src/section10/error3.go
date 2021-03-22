//go 에러 처리 기초(2)

package main

import (
	"fmt"
	"errors"
)


func main(){
	//에러 처리
	//errors 패키지의 new 메소드를 활용한 에러 생성
	//errorf 보다 더 많이 사용
	var err1 error = errors.New("Error occured - 1")
	err2 := errors.New("Error occured - 2")

	fmt.Println(err1)
	fmt.Println(err1.Error())

	fmt.Println(err2)
	fmt.Println(err2.Error())
}