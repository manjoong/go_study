// 함수 기초(1)
package main

import "fmt"

//함수 선언 위치는 어느 곳이든 가능

func sum(i int, f func(int, int)){
	f(i, 10)
}

func add(a, b int){
	fmt.Println(a+b)
}
///////

func multi_value(i int){
	i = i*10
}

func multi_reference(x *int){
	*x *=10

}

func main(){
	sum(100, add)

	a:=100
	multi_reference(&a)

	fmt.Println(a)

	//reference value


	

}