package main

import "fmt"


func multi_fly(x int, y int)(r1 int, r2 int){

	r1 = x *10
	r2 = y*20

	return r1, r2
}

func main(){
	a, b := multi_fly(10, 10)

	fmt.Println(a, b)
}