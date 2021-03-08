package main

import (
	"fmt"

	"sync"

)



func main(){
	//고루틴 동기화 고급
	//대기 그룹

	//실행 흐름 변경(고루틴이 졸료 될 때 까지 대기 가능)
	//waitgroup : add(고루틴 추가), done(작업 종료 알림), wait(고루틴 종료시까지 대기)
	//Add로 추가 된 고루틴 개수와 done으로 종료되는 알림 횟수가 같아야 정확하게 동작한다.(add==done)


	wg := new(sync.WaitGroup)

	for i := 0; i< 100; i++{
		wg.Add(1)
		go func(n int){
			fmt.Println("waitGroup: ", n)
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println("waitGroup end ")
}