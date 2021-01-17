package main

import "fmt"

func main(){

	arr1 := [5]int{1, 10, 100, 1000, 10000}

	for i:=0; i< len(arr1); i++{
		fmt.Printf("%d\n", arr1[i]) 
	}

	//가장 많이 사용
	arr2 := [5]int{7,77,777,77777,777777}

	for i, v:=range arr2{
		fmt.Println(i, v)
	}

	//인덱스 생략
	for _, v:=range arr2{
		fmt.Println(v)
	}

	for i:=range arr2{
		fmt.Println(i)
	}
}