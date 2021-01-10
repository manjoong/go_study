package main

import (
	"fmt"
	"math"
)

func main() {
	//숫자 연산(산술, 비교)
	//타입이 같아야 가능(엄격)
	//다른 타입끼리는 반드시 형 변환 후 연산
	// +, -, *, %, / , <<, >>, &, ^

	var n1 uint8 = math.MaxUint8
	var n2 uint16 = math.MaxUint16
	var n3 uint32 = math.MaxUint32
	var n4 uint64 = math.MaxUint64

	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)
	fmt.Println(n4)

}
