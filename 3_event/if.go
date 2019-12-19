package main

import "fmt"


func main(){
	for i :=1;i<10;i++{
		if i%2 ==0 {
			fmt.Println("i는 짝수 입니다.")
		}else if i%2 == 1 {
			fmt.Println("i는 홀수 입니다.")
		}else{
			fmt.Println("이런 수는 나올수가 없습니다.")
		}
	}
}