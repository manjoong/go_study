//고루틴 동기화 기초(5)

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main(){
	//고루틴 동기화 객체
	//동기화 상태(조건) 메소드 사용
	//wait, notifyAll : 기타언어
	//wait, signal, broadcast : 

	runtime.GOMAXPROCS(runtime.NumCPU())
	
	var mutex = new(sync.Mutex)
	var condition = sync.NewCond(mutex)

	c := make(chan int, 5) //비동기 버퍼 채널

	for i :=0; i<5; i++{
		go func(n int){
			mutex.Lock()
			c <- 777
			fmt.Println("Goroutine Waiting", n)
			condition.Wait()
			fmt.Println("Waiting End: ", n)
			mutex.Unlock()
		
		}(i)
	}

	for i :=0 ; i< 5; i++{
		<-c
	}
}