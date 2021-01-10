package main

import "fmt"

func main(){
	for i :=0; i<5; i++{
		fmt.Println(i)
	}

	// for {
	// 	fmt.Println("실행하면 큰일나")
	// }

	// loc := []string{"seoul","boosan","incheon"}

	sum, i := 0,0

	for {
		if i>100{
			break
		}
		sum += i
		i++
	}
	fmt.Println(sum)

	for i, j:=0,0; i<=10; i,j=i+1,j+1{
		fmt.Println(i, j)
	}

}