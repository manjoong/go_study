package main

import "fmt"

//값을 추가할 수 있는 append
func main() {
	//s4가 s3를 재할당하는 컨벤션을 따르기 때문에 전달된 슬라이스의 내부배열을 참조하여 슬라이스를 선언하면 변수값이 같이 바뀐다.
	s1 := []string{"s1", "s1"}
	s2 := append(s1, "s2", "s2")
	s3 := append(s2, "s3", "s3")
	s4 := append(s3, "s4", "s4")
	s4[0] = "X"
	fmt.Println(s1, s2, s3, s4)

	//따라서 동일한 변수에 재할당하는 게 일반적이다.
	s5 := []string{"s5", "s5"}
	s5 = append(s5, "s6", "s6")
	s5 = append(s5, "s7", "s7")
	s5 = append(s5, "s8", "s8")
	fmt.Println(s5)


}