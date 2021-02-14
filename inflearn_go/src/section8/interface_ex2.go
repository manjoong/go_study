//인터페이스 고급
package main

import "fmt"
import "reflect"

func main()  {
	//타입 변환
	//실행(런타임) 시에는 인터페이스에 할당한 ㅂㄴ수는 실제 타입으로 변환 후 사용해야 하는 경우
	//인터페이스(타입) -> 형변환
	//

	//예제1
	var a interface{} = 15
	b:= a
	c := a.(int)
	// d := a.(float64)
	

	fmt.Println(a)
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(b)
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(c)
	fmt.Println(reflect.TypeOf(c))


	//예제2 (저장된 실제 타입 검사)

	if v, ok := a.(int); ok{
		fmt.Println(v, ok)
	}

}

