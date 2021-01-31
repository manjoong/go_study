package main	

import "fmt"

func main()  {
	//익명 함수 선언
	//예제1 type 구조체명 타입

	car1 := struct{name, color string}{"520d", "red"}

	fmt.Println(car1)

	//예제 2
	cars := []struct {name, color string}{{"520d", "red"},{"530d", "red"},{"540d", "red"}}


	for _, c:= range cars{
		fmt.Printf("(%s, %s)-----%#v\n", c.name, c.color, c)
	}

	
}