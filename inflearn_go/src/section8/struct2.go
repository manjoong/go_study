package main

import "fmt"

type Account struct {
	number string
	balance float64
	interest float64
}
//리시버를 통해서 구조체<->메서드를 연결

func (a Account) Calculate() float64{
	return a.balance + (a.balance* a.interest)
}

func  main()  {
	//다양한 선언 방법
	//&struct, struct: &struct 포인터를 받아오고 역참조를 또 하기 때문에 속도는 조금 느리다
	//인터페이스 메소드를 선언만 해둔 후-> 오버라이딩 해서 메소드에 포인터 리시버를 사용할 경우 반드시 &struct로 넘겨야 함

	//선언 방법1
	var kim *Account = new(Account)

	kim.number = "245-901"
	kim.balance = 10000
	kim.interest = 0.015

	//선언 방법2
	hong := &Account{number:"253-2333", balance: 1233, interest:0.02}

	//선언 방법3
	lee:= new(Account)
	lee.number="123322"
	lee.balance=1233223
	lee.interest=0.024


	// kim := Account{number: "245-601", balance: 10000000, interest: 0.015}
	// lee := Account{number: "245-421", balance: 12000000}
	// park := Account{number: "245-903", interest: 0.025}
	// cho := Account{"245-903", 150000, 0.03}

	fmt.Println(kim)
	fmt.Println(hong)
	fmt.Println(lee)

	fmt.Printf("%#v", kim)
	fmt.Printf("%#v", hong)
	fmt.Printf("%#v", lee)

	// //예제2

	// fmt.Println(kim.Calculate)

	//예제2
	fmt.Println(int(kim.Calculate()))

	
}