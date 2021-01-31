//구조체 심화
//상속이란 개념으로 생각하면됨. 사실은 메서드 재사용 
package main
import "fmt"

type Employee struct{
	name string
	salary float64
	bonus float64

}

func (e Employee) Calculate() float64  {
	return e.salary + e.bonus
}

type Executives struct {
	Employee //is a 관계
	specialBonus float64
}

func main()  {
	//구조체 임베디드 패턴
	//다른 관점으로 메소드를 재사용하는 장점 제공
	//상속을 허용하지 않는 go 언어에서 메소드 상속 활용을 위한 패턴

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
	fmt.Println(int(ex1.Calculate()+ex1.specialBonus))

	
}