package main

import(
	"fmt"
)
//맵은 다중 리턴(값, bool)값을 가지고 있으며 해당 키가 없을 경우엔 bool값이 false로 들어온다.
func main(){
	counter := map[string]int{"a" : 1, "b" : 2}
	var value int
	var ok bool

	value, ok = counter["a"]
	fmt.Println(value, ok)
	value, ok = counter["c"]
	fmt.Println(value, ok)

	//delete를 통해 맵의 키를 삭제할 수 있다.
	value, ok = counter["b"]
	fmt.Println(value, ok)
	delete(counter, "b")
	value, ok = counter["b"]
	fmt.Println(value, ok)
}