// 함수 기초(1)
package main

import "fmt"
import "strconv" 
//함수 선언 위치는 어느 곳이든 가능

func helloGolang(){
	fmt.Println("ex1 : Hello Golang!!")
}

func sayOne(m string){
	fmt.Println(m)
}

func add(x int, y int) int{
	fmt.Println(x+y)
	return(x+y)
}
func main(){
	//함수 선언
	//fumc 키워드로 선언
	//func 함수명 (매개 변수) -> 반환 타입 or 반환값
	//func 함수명()
	//func 함수명(매개 변수)
	//func 함수명() 매개변수 없음, 반환값 없음
	//타 언어와 달리 반환값 (return value) 다수 가능 -> 다중 반환 함수

	helloGolang()
	sayOne("hi")
	add(5, 5)
	fmt.Println(strconv.Itoa(add(5,5)))

}