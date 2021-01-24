// 함수 기초(1)
package main

import "fmt"

//함수 선언 위치는 어느 곳이든 가능

func multi_fly(x int, y int)(int, int){
	return x *10, y *10
}

func array_multi(a,b,c,d,e int) (int, int, int, int, int){
	return a*1, b*2, c*3, d*4, e*5
}

func main(){

	a, b := multi_fly(10, 5)
	fmt.Println(a, b)


	c, _,e,f,g := array_multi(1,2,3,4,5)

	fmt.Println(c, e,f,g )



}