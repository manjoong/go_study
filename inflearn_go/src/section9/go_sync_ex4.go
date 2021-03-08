package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"

)


//원자성 사용 안할때
func main(){
	//고루틴 동기화 고급
	//원자성 사용 -> 기능적으로 분할 불가능한 완전 보증된 일련의 조작, 모두 성공 or 모두 실패
	//모든 조작이 완료 될 때까지 다른 프로세스 개입 불가
	//sync/atomic 에서 원자적 연산자 제공

	runtime.GOMAXPROCS(runtime.NumCPU())

	var cnt int64 = 0
	wg := new(sync.WaitGroup)

	for i := 0; i< 5000; i++{
		wg.Add(1)
		go func(n int){
			//cnt +=1
			atomic.AddInt64(&cnt, 1)
			wg.Done()
		}(i)
	}

	for i:=0; i< 2000; i++ {
		wg.Add(1)
		go func(n int){
			//cnt -=1
			atomic.AddInt64(&cnt, -1)
			wg.Done()
		}(i)
	}

	//add (7000) == done(7000)
	wg.Wait()

	finalCnt := atomic.LoadInt64(&cnt)
	fmt.Println("waitGroup end ", finalCnt)
}