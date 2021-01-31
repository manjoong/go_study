//구조체 심화

package main
import "fmt"

type Account struct{
	number string
	balance float64
	interest float64
}

func NewAccount(number string, balance float64, interest float64) *Account { //포인터 반환 아닌 경우 값 복사
	return &Account{number, balance, interest} //구조체 인스턴를 생성 한 뒤 포인터 리턴
}

func main()  {
	//예제1
	kim := Account{number: "245-1233", balance:10000, interest: 0.013}

	var lee *Account = new(Account)
	lee.number= "245-023" //getter, setter
	lee.balance=3111111
	lee.interest=0.022

	fmt.Println(kim)
	fmt.Println(lee)

	//예제2
	park := NewAccount("24244-1244", 173333, 0.04)
	fmt.Println(park)


	
}