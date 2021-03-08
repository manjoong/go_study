//고루틴 동기화 기초

package main

import (
	"fmt"
	"time"
)


func main(){
	//뮤텍스 : 상호 배제 -> 뜨레드(고루틴)들이 서로 runningtime에 서로 영향을 주지 않게 단독으로 실행 되게 하는 기술
	//false
	//뮤텍스 : 여러 고루틴에서 작업하는 공유 데이터 보호

	//동기화 사용하지 않는 경우 예제
	//쓰기, 읽기 동작 순서가 일정하지 않아 잘못된 오류를 반환 할 가능성 증가


	data := 0
	go func(){
		for i := 0; i<=10; i++{
			data+=1
			fmt.Println("Write: ", data)
			time.Sleep(200*time.Millisecond)
		}
	}()
	
	go func(){
		for i := 0; i<=10; i++{
			fmt.Println("Read: ", data)
			time.Sleep(1*time.Second)
		}
	}()

	time.Sleep(5*time.Second)
	
	

}