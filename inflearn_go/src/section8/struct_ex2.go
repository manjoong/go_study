//구조체 심화

package main
import "fmt"

type Account struct{
	number string
	balance float64
	interest float64
}

func CalculateD(a Account){
	a.balance = a.balance + (a.balance * a.interest)
}

func CalculateP(a *Account){
	a.balance = a.balance + (a.balance * a.interest)
}

func main()  {
	//예제1
	kim := Account{"245-901", 900000, 0.015}
	lee := Account{"23311-13", 8222, 0.3}

	fmt.Println(kim)
	fmt.Println(lee)

	CalculateD(kim)
	CalculateP(&lee)

	fmt.Println(int(kim.balance))
	fmt.Println(int(lee.balance))
	
}