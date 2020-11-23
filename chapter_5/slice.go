package main

import "fmt"

//슬라이스는 make연산자를 이용해 미리 선언해서 만들거나 
// 미리 값을 알고 있는 경우, 리터럴로 만들 수 있다.
func main() {
	numbers := make([]float64, 3)
	numbers[0] = 19.7
	numbers[2] = 25.2
	for i, number := range numbers {
		fmt.Println(i, number)
	}
	var letters = []string{}
	letters = append(letters, "a")
	fmt.Println(len(letters))
	for i, letter := range letters{
		fmt.Println(i, letter)
	}
}