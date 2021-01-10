package main

import "fmt"

func main(){
	a := 30 /15

	switch a{
	case 2, 4, 6:
		fmt.Println("a -> ", a,"는 짝수")
	case 1, 3, 5, 7:
		fmt.Println("a -> ", a,"는 홀수")
	}

	switch e:= "python"; e{
	case "java":
		fmt.Println("java!")
		fallthrough
	case "python":
		fmt.Println("python!")
		fallthrough
	case "c":
		fmt.Println("c!")
		fallthrough
	case "go":
		fmt.Println("go!")
	}
}