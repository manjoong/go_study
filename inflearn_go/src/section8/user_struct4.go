package main

import "fmt"

type shoppingBasket struct{cnt, price int}

func (b shoppingBasket) purchase() int{
	return b.cnt * b.price
}

//원본이 수정(참조 형식)
func (b *shoppingBasket) rePurchaseP(cnt, price int){
	b.cnt += cnt
	b.price += price
}

//원본 수정 x (값 전달 방식)
func (b shoppingBasket) rePurchaseD(cnt, price int){
	b.cnt += cnt
	b.price += price
}

func main(){

	//리시버(구조체 메소드), 전달(값 참조)형식
	//함수는 기본적으로 값 호출 -> 변수의 값 복사 -> 전달 방식
	//리시버(구조체) 도 마찬가지로 포인터를 활용해서 메소드 내에서 원본 수정 가능

	//예제1
	bs1:= shoppingBasket{3, 5000}
	fmt.Println(bs1.purchase())

	//원본값이 수정
	bs1.rePurchaseP(7, 5000) //보통 포인터 형의 매개변수를 받는 함수는 변수 전달시 &기호를 넣어야 하지만 구조체는 알아서 들어감
	fmt.Println(bs1.purchase())

	//수정x
	bs1.rePurchaseD(10, 0)
	fmt.Println(bs1.purchase())



}