//자료형: 포인터(1)

package main

import "fmt"

func main()  {
	// 예제1
	i := 7 
	p := &i

	fmt.Println(i)
	fmt.Println(*p)
	fmt.Println(&i)
	fmt.Println(p)
	
	*p++
	fmt.Println(i)
	fmt.Println(*p)

	*p = 7777
	fmt.Println(i)
	fmt.Println(*p)

	i = 77
	fmt.Println(i)
	fmt.Println(*p)
	fmt.Println(&i)
	fmt.Println(p)






}