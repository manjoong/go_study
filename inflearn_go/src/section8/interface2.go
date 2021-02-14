//인터페이스 기본

package main


import "fmt"

type Dog struct{
	name string
	weight int
}

//bite메소드 구현

func (d Dog) bite(){
	fmt.Println(d.name, "bite!")
}

//동물의 행동 인터페이스 선언
type Behavior interface{
	bite()
}

func main()  {
	//인터페이스 구현 예제
	dog1:=Dog{"poll", 10}
	var interface1 Behavior
	interface1 = dog1
	interface1.bite()

	//예제2
	dog2:=Dog{"marry", 12}
	inter2 := Behavior(dog2)

	inter2.bite()

	inters := []Behavior{dog1, dog2}

	//인덱스 형태로 실행

	for idx, _ := range inters{
		inters[idx].bite()
	}

	//값 형태로 실행(인터페이스)

	for _, val := range inters {
		val.bite()
	}

}