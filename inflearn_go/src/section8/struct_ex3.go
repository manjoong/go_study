//구조체 심화

package main
import "fmt"

type Account struct{
	number string
	balance float64
	interest float64
}
//리시버
func (a Account) CalculateD(bonus float64){
	a.balance = a.balance + (a.balance * a.interest)+ bonus
}

func (a *Account) CalculateF(bonus float64){
	a.balance = a.balance + (a.balance * a.interest)+ bonus
}



func main()  {
	//예제1

	//정의 : 구조체 인스턴스 값 변경시 -> 포인터를 전달, 보통의 경우 값전달
	kim := Account{"245-901", 900000, 0.015}
	lee := Account{"23311-13", 8222, 0.3}

	fmt.Println(kim)
	fmt.Println(lee)

	

	lee.CalculateD(100000)
	kim.CalculateF(100000)
	fmt.Println(int(lee.balance))
	fmt.Println(int(kim.balance))
	
}