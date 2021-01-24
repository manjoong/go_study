// 함수 기초(1)
package main

import "fmt"
//함수 선언 위치는 어느 곳이든 가능




func main(){
	//익명 함수 (아나니머스 펑션)
	//즉시실행(선언과 동시에)

	func (){
		fmt.Println("익명함수!")

	}()

	msg:= "안녕!"
	func(m string){
		fmt.Println(m)
	}(msg)

	func(x, y int){
		fmt.Println(x+y)
	}(10, 20)

	r:= func(x, y int) int {
		return x*y
	}(10, 100)

	fmt.Println(r)
}