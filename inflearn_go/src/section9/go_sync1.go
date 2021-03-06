//고루틴 동기화 기초

package main

import (
	"fmt"
	"runtime"
)
//구조체 선언(공유 데이터)
type count struct {
	num int
}

func (c *count) increment(){
	c.num += 1
}

func (c *count) result(){
	fmt.Println(c.num)
}

func main(){
	//고루틴 동기화 예제
	//실행 흐름 제어 및 변수 동기화 가능
	//공유 데이터 보호가 가장 중요

	//동기화 전체 cpu tkdyd

	runtime.GOMAXPROCS(runtime.NumCPU())
	c := count{num : 0}
	done := make(chan bool)
	for i:=1; i<= 10000; i++{
		go func(){
			c.increment()
			done <- true
			runtime.Gosched()//cpu 양보 해주는거
		}()
	}

	for i := 1; i<=10000; i++{
		<-done
	}

	c.result()

}