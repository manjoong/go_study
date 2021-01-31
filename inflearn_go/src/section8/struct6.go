package main	

import "fmt"


type car struct{
	name string "차량명"
	color string "색상"
	company string "제조사"
	detail spec "상세"
}

type spec struct{
	length int "전장"
	height int "전고"
	width int "전축"
}

func main()  {
	//중첩 구조체
	//예제 1
	car1 := car{
		"520d",
		"silver",
		"bmw",
		spec{4000, 1000, 2000},

	}

	//내부 spec구조체 값 출력
	fmt.Println(car1.name)
	fmt.Println(car1.color)
	fmt.Println(car1.company)
	fmt.Printf("%#v\n", car1.detail)

	//내부 spec구조체 필드값
	fmt.Println(car1.detail.length)
	fmt.Println(car1.detail.width)
	fmt.Println(car1.detail.height)
	// fmt.Printf("%#v\n", car1.detail)
}
