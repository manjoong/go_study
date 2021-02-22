package main

import (
	"fmt"
	"runtime"
)

func main(){
	//예제1(비동기: 버퍼 사용)
	//버퍼: 발신 -> 가득차면 대기, 비어있으면 작동, 수신 -> 비어있으면 대기, 가득차 있으면 작동

	runtime.GOMAXPROCS(1)
	ch := make(chan bool, 2)   //채널의 버퍼의 용량
	cnt := 12

	go func(){ //고루틴으로 쓸드로 실행 하기 때문에 수신 할때 까지 대기 함 (수신과 송신이 동기식으로 실행됨)
		for i := 0; i<cnt; i++{
			ch <- true
			fmt.Println("GO",i)
		}
	}()

	for i :=0; i<cnt;i++{
		<-ch
		fmt.Println("MAIN", i)
	}
}