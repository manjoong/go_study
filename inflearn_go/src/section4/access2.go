//패키지 접근제어

package main

import (
	"fmt"
	_ "section4/lib"        //사용하지 않아도 저장시 없어지지 않도록
	checkUp "section4/lib2" //식별자를 통해 패키지를 사용
)

func main() {
	fmt.Println("10보다 큰 수?: ", checkUp.CheckNum(11))
}
