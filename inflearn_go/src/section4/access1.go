//패키지 접근제어

package main

import (
	"fmt"
	"section4/lib2"
)

func main() {
	fmt.Println("100보다 큰 수?:", lib2.CheckNum1(101))
	fmt.Println("1000보다 큰 수?:", lib2.CheckNum2(999))
}
