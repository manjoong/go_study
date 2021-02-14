//인터페이스 기본

package main


import "fmt"

type Dog struct{
	name string
	weight int
}

type Cat struct{
	name string
	weight int
}

//구조체 dog메소드 구형

func (d Dog) bite(){
	fmt.Println(d.name, "Dog bite!")
}

func (d Dog) sound(){
	fmt.Println(d.name, "Dog barks!")
}

func (d Dog) run(){
	fmt.Println(d.name, "Dog runnning!")
}

func (d Cat) bite(){
	fmt.Println(d.name, "Cat craw!!")
}

func (d Cat) sound(){
	fmt.Println(d.name, "Cat cries!")
}

func (d Cat) run(){
	fmt.Println(d.name, "Cat runnning!")
}

//동물의 행동 인터페이스 선언 메서드만을 가지고 타입을 판단하는것 -> 덕타입(메소드로만 타입을 판단)
type Behavior interface{
	bite()
	sound()
	run()
}

//인터페이스를 파라미터 받는다.
func act(animal Behavior){
	animal.bite()
	animal.sound()
	animal.run()
}

func main()  {
	//인터페이스 구현 예제
	dog1:=Dog{"poll", 10}
	cat1:=Cat{"nami", 5}

	act(dog1)
	act(cat1)

	// var interface1 Behavior
	// interface1 = dog1
	// interface1.bite()

	

}