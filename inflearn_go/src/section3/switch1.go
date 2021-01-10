//switch1
package main

import "fmt"

func main(){
	//제어문 조건문
	//switch

	//type으로 분기 가능
	a := -7
	switch {
	case a < 0:
		fmt.Println(a, "는 음수")
	case a > 0:
		fmt.Println(a, "는 양수")
	case a == 0:
		fmt.Println(a, "는 0")
	}

	switch b := 27; {
	case b < 0:
		fmt.Println(b, "는 음수")
	case b > 0:
		fmt.Println(b, "는 양수")
	case b == 0:
		fmt.Println(b, "는 0") 
	}

	switch  c:="go"; c{
	case "go":
		fmt.Println("Go!!")
	case "java":
		fmt.Println("java!!")
	default:
		fmt.Println("일치하는 값 없음")
		
	}

	switch i, j :=20, 30;{
	case i < j:
		fmt.Println("i는 j보다 작다")
	case i > j:
		fmt.Println("i는 j보다 크다")
	case i == j:
		fmt.Println("i는 j와 같다")

	}
	
}