package main

import "fmt"

func main(){
	
	i := 0
	for i<=4 {
		fmt.Println(i)
		i = i+1
	}

	for j := i; j<=7; j++ {
		if j == i {
			fmt.Println("두번째 루프")
			continue
		}
		fmt.Println(j)
	}


}