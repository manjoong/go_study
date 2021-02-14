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


func (d Dog) run(){
	fmt.Println(d.name, "Dog runnning!")
}


func (d Cat) run(){
	fmt.Println(d.name, "Cat runnning!")
}



//인터페이스를 파라미터 받는다.
func act(animal interface{ run() }){ //익명 인터페이스 (타입 정의 x)
	animal.run()
}

func main()  {

	//익명 인텊이스 사용 예제(즉시 선언 후 사용)

	
	dog1:=Dog{"poll", 10}
	cat1:=Cat{"nami", 5}

	act(dog1)
	act(cat1)

	// var interface1 Behavior
	// interface1 = dog1
	// interface1.bite()

	

}