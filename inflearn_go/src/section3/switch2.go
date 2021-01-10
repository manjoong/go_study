//switch1
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	//제어문 조건문
	//switch

	//type으로 분기 가능
	// a := -7

	rand.Seed(time.Now().UnixNano())
	switch i:=rand.Intn(100); {
	case i >= 50 && i<100:
		fmt.Println("i -> ", i, " 50이상 100미만이다.")
	case i >= 25 && i<60:
		fmt.Println("i -> ", i, " 50이하 이다.")
	default:
		fmt.Println("i -> ", i, " 50이상 100미만이다.")
		
	}

	
}