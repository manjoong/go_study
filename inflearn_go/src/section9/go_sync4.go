//고루틴 동기화 기초

package main

import (
	"fmt"
	"time"
	"runtime"
	"sync"
)


func main(){
	//뮤텍스 : 상호 배제 -> 뜨레드(고루틴)들이 서로 runningtime에 서로 영향을 주지 않게 단독으로 실행 되게 하는 기술
	//false
	//뮤텍스 : 여러 고루틴에서 작업하는 공유 데이터 보호
	//RWMutex : 쓰기 Lock ->쓰기 시도 중에는 다른 곳에서 이전 값을 읽으면 x, 읽기 락, 쓰기 락 전부 방어
	//RMutex : 읽기 lock -> 읽기 시도 중에 값이 변경 방지 즉, 쓰기 락 방어(방지)


	//동기화 사용하지 않는 경우 예제
	//쓰기, 읽기 동작 순서가 일정하지 않아 잘못된 오류를 반환 할 가능성 증가

	runtime.GOMAXPROCS(runtime.NumCPU())


	data := 0
	mutex := new(sync.RWMutex) 


	go func(){
		for i := 0; i<=10; i++{
			//쓰기 뮤텍스 잠금
			mutex.Lock()
			data+=1
			fmt.Println("Write: ", data)
			time.Sleep(200*time.Millisecond)
			mutex.Unlock()
		}
	}()

	go func(){
		for i := 0; i<=10; i++{
			//읽기 뮤텍스 잠금
			mutex.RLock()
			fmt.Println("Read: ", data)
			time.Sleep(1*time.Second)
			//읽기 뮤텍스 해제
			mutex.RUnlock()
		}
	}()
	
	go func(){
		for i := 0; i<=10; i++{
			//읽기 뮤텍스 잠금
			mutex.RLock()
			fmt.Println("Read2: ", data)
			time.Sleep(1*time.Second)
			//읽기 뮤텍스 해제
			mutex.RUnlock()
		}
	}()

	time.Sleep(5*time.Second)
	
	

}