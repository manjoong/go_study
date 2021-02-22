package main

import (
	"fmt"
	"time"
)

func main(){ 
	//채널
	//예제1

	ch1 := make(chan int)
	ch2 := make(chan string)

	go func(){
		for{
			num:= <-ch1 //값 수신
			fmt.Println(num)
			time.Sleep(250 * time.Millisecond)
		}
	}()

	go func(){
		for{
			ch2 <- "Golang HI" 
			time.Sleep(500 * time.Millisecond)
		}
	}()

	go func(){
		for {
			select{
			case ch1 <- 777: //발신 용도123
			case str := <-ch2:
				fmt.Println("ch2:", str)
			}
		}

	}()
	time.Sleep(7 * time.Second)
	
}