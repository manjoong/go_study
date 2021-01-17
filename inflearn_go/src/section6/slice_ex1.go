//자료형 슬라이스 심화
package main

import "fmt"

func main()  {
	s1 := []int{1,2,3,4,5} //길이가 5개고 용량이 10일때 어팬드를 쓰면 7,10이 됨 만약 길이가 10을 넘어간다면? 
	//지금은 그때 그때 마다 용량이 변경 되어진 슬라이스를 재할당 해주는 방식. 하지만 사용자가 많아지면? 재할당에 리소스가 소모됨. 따라서 용량을 미리 선언하는 것이 효율적

	s1 = append(s1, 6, 7 ,8 ,9, 1, 2, 3)

	s2 := []int{8, 9, 10, 11,12}
	s3 := []int{13, 14, 15, 16,17}

	s1 = append(s1, 6, 7 ,8 ,9, 1, 2, 3)
	s2 = append(s1, s2...) //슬라이스 뒤에 슬라이스를 삽입할땐 ...
	s3 = append(s2, s3[0:3]...) //추출 후 병합

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)

	//예제 2
	s4 := make([]int, 0, 5)
	for i :=0; i<15; i++{
		s4= append(s4, i)
		fmt.Println(len(s4), cap(s4), s4)
	}
}