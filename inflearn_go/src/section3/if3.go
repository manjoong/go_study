package main

import "fmt"

func main(){

	b := 110

	//예제 1
	if b >=120{
		fmt.Println("120이상")
	}else if b >= 100 && b<120{
		fmt.Println("100이상 120이하")
	}
}