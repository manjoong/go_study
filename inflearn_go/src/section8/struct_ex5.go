//구조체 심화

//오버 라이딩 = 어떤 메소드를 수정하여 사용 하는것
package main
import "fmt"

type Employee struct{
	name string
	salary float64
	bonus float64

}

func (e Employee) Calculate() float64  { //=> 기존 사용되던 이 메서드를
	return e.salary + e.bonus
}

func (e Executives) Calculate() float64  { //=> 이렇게 오버라이딩 하여 사용함
	return e.salary + e.bonus + e.specialBonus
}

type Executives struct {
	Employee //is a 관계
	specialBonus float64
}

func main()  {
	//구조체 임베디드 메소드 오버라이딩

	//직원
	ep1:= Employee{"kim", 200000, 150000}
	ep2:= Employee{"park", 200000, 30000}

	//임원
	ex1 := Executives{
		Employee{"lee", 500000, 100000},
		1000000,
	}

	fmt.Println(int(ep1.Calculate()))
	fmt.Println(int(ep2.Calculate()))

	//employee구조체를 통해서 메소드 호출
	fmt.Println(int(ex1.Calculate()))
	fmt.Println(int(ex1.Employee.Calculate()+ex1.specialBonus))

	
}