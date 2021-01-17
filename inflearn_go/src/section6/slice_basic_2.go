//자료형 배열

package main

import "fmt"

func main(){
	//슬라이스(슬라이스 참조 타입 증명)

	arr1 := [3]int{1,2,3}
	var arr2 [3]int

	arr2 = arr1 //값을 복사
	arr2[0] = 7

	fmt.Println(arr1)
	fmt.Println(arr2)



	//예제2 (슬라이스)
	slice1 := []int{1,2,3}
	var slice2 []int

	slice2 = slice1 //참조 위치를 복사
	slice2[0] = 7
	fmt.Println(slice1)
	fmt.Println(slice2)


	//예제3
	slice3 := make([]int, 50, 100)
	fmt.Println(slice3[4])
	fmt.Println(slice3[49])

	//예제 4
	for i, v := range slice1{
		fmt.Println(i, v)
	}

	// var slice1 = []int{}
	// slice2 := []int{}
	// slice3 := []int{1,2,3,4,5}
	// slice4 := [][]int{
	// 	{1,2,3,4,5},
	// 	{6,7,8,9,10},
	// }
	// // slice2[0] = 1
	// slice3[4]=10

	// fmt.Printf("%-5T %d %d %v\n", slice1, len(slice1), cap(slice1), slice1)
	// fmt.Printf("%-5T %d %d %v\n", slice2, len(slice2), cap(slice2), slice2)
	// fmt.Printf("%-5T %d %d %v\n", slice3, len(slice3), cap(slice3), slice3)
	// fmt.Printf("%-5T %d %d %v\n", slice4, len(slice4), cap(slice4), slice4)

	// //예제 2
	// var slice5 []int = make([]int, 5, 10) //용량에 근접하게 잡을수 있는게 가장 예측하여 사용할 수 있기 때문에 size를 특정한다.그냥 사이즈를 빼면 용량도 5로 선언
	// var slice6 = make([]int, 5)
	// slice7 :=make([]int, 5, 100)
	// slice8 :=make([]int, 5)

	// slice6[2] = 7 //make로 만들면 0으로 초기화 해주기 때문에 해당 명령어가 먹음

	// fmt.Printf("%-5T %d %d %v\n", slice5, len(slice5), cap(slice5), slice5)
	// fmt.Printf("%-5T %d %d %v\n", slice6, len(slice6), cap(slice6), slice6)
	// fmt.Printf("%-5T %d %d %v\n", slice7, len(slice7), cap(slice7), slice7)
	// fmt.Printf("%-5T %d %d %v\n", slice8, len(slice8), cap(slice8), slice8)
	// //그냥 배열처럼 선언할때랑 make로 처리할때 다른점 => make는 0으로 초기화 해준다.

	// //예제3
	// var slice9 []int //nil 슬라이스(길이와 용량이 0)
	// if slice9 == nil{
	// 	fmt.Println("slice9은 nil")
	// }

}