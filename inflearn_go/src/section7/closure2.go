package main

import "fmt"

func main()  {
	//예제 1
	cmt := increaseCnt()

	fmt.Println(cmt())
	fmt.Println(cmt())
	fmt.Println(cmt())
	fmt.Println(cmt())
	fmt.Println(cmt())


}

func increaseCnt() func() int{
	n :=0 //지역변수 캡쳐 -

	return func() int{
		n += 1
		return n
	}
}