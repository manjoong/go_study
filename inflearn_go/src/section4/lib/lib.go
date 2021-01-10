package lib //패키지 명은 폴더 이름

import "fmt"

func init() {
	fmt.Println("lib init start!")
}

func CheckNum(c int32) bool {
	return c > 10
}
