//자료형: 포인터(1)

package main

import "fmt"

func main()  {
	//변수의 지역성, 연속된 메모리의 참조

	//파이썬, 자바 -> 포인터 없음~~
	//포인터를 지원 x(파이썬, java, c#)
	//주소의 값은 직접 변경 불가능(잘못된 코딩으로 인한 버그 방지)
	//*로 사용
	//nil로 초기화

	//예제1
	var a *int //방법 nil로 초기화
	var b *int = new(int) //정석 주소값으로 초기화

	fmt.Println(a)
	fmt.Println(b)

	i := 7 

	a= &i
	b = &i
	*a = 77

	fmt.Println(i, &i)
	fmt.Println("")
	fmt.Println(a)
	fmt.Println(&a)
	fmt.Println(*a) //역참조 역으로 번지수를 찾아가 값을 보는 것
	fmt.Println("")
	fmt.Println(b)
	fmt.Println(&b)
	fmt.Println(*b)

	var c = &i
	d := &i

	

	fmt.Println(c)
	fmt.Println(&c)
	fmt.Println(*c) //역참조 역으로 번지수를 찾아가 값을 보는 것
	fmt.Println("")
	fmt.Println(d)
	fmt.Println(&d)
	fmt.Println(*d)



}