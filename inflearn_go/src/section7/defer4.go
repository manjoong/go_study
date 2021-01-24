package main

import "fmt"

func start(t string){
	fmt.Println("start", t)
}

func end(t string){

	fmt.Println("end", t)
}

func a(){
	defer end(start("b"))
	fmt.Println("in a")
}

func main()  {
	a()
	
}