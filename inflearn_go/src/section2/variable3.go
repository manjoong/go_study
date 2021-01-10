package main


import "fmt"

func main(){
	//짧은 선언 
	//특징: 반드시 함수 안에서만 사용.... 전역으론 사용불가. 선언 후, 재할당시 예외 발생(이 안에서만 사용할꺼야. 일회성으로만 선언할꺼야)
	//주로 제한된 범위의 함수 내에서 사용될 경우.. 코드 가동성을 높임.
	shortVar1 := 3
	shortVar2 := "test"
	shortVar3 := false
	fmt.Println("shortVar1: ", shortVar1)
	fmt.Println("shortVar2: ", shortVar2)
	fmt.Println("shortVar3: ", shortVar3)

	if i:=10; i< 11{
		fmt.Println("short variable test success")
	}
}
