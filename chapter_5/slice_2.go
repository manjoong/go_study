package main

import "fmt"

//슬라이스는 배열 데이터에 대한 '뷰' 일 뿐이다.(기본적으로 배열임)
//따라서 배열을 참조하고 있는 슬라이스는, 배열의 값을 변경하면 슬라이스도 변경 된다.
func main() {
	array1 := [5]string{"a","b","c","d","e"}
	slice1 := array1[0:3]
	array1[1] = "X"
	fmt.Println(array1)
	fmt.Println(slice1)
}

//따라서 배열을 먼저 만들고 슬라이스로 참조하기 보다는, make나 슬라이스 리터럴로 만드는게 더 낫다.