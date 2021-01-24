// 함수 기초(1)
package main

import "fmt"
//함수 선언 위치는 어느 곳이든 가능




func multi_fly(x, y int) (r int){
	r = x *y

	return r
}

func add_fly(x, y int) (r int){
	r = x + y

	return r
}
func main(){
	//함수 고급
	//함수를 슬라이스에 할당

	f := []func(int, int) int{multi_fly, add_fly}
	a:= f[0](10, 10)
	b:= f[1](10, 10)
	fmt.Println(a, b)
	//함수를 변수에 할당

	var f1 func(int, int) int = multi_fly //한번 찾아봐야함
	f2 := add_fly //오 sum이 뭐야
	fmt.Println(f1(20, 20))
	fmt.Println(f2(32, 1))


	m := map[string] func(int, int) int{
		"multi": multi_fly,
		"sum": add_fly,
	}

	fmt.Println(m["multi"](30, 30))


	 

}