//문자열 결합

package main

import (
	"fmt"
	"strings"
)

func main() {
	//일반 결합 연산

	var str1 string = "Hello " +
		"world "
	var str2 string = "manjoong"

	fmt.Println(str1 + str2)

	//join
	strSet := []string{}
	strSet = append(strSet, str1)
	strSet = append(strSet, str2)

	fmt.Println(strings.Join(strSet, "-----"))

}
