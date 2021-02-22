package main

import (
	"fmt"
	"time"
)

func main(){
	//예제1(동기: 버퍼 미상)

	ch := make(chan bool)
	cnt := 6

	go func(){ //고루틴으로 쓸드로 실행 하기 때문에 수신 할때 까지 대기 함 (수신과 송신이 동기식으로 실행됨)
		for i := 0; i<cnt; i++{
			ch <- true
			fmt.Println("GO",i)
			time.Sleep(1 * time.Second)
		}
	}()

	for i :=0; i<cnt;i++{
		<-ch
		fmt.Println("MAIN", i)
	}
}