package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	//문자열
	//큰타옴표 "", 백스쿼트 ₩₩
	//char타입 존재x rune(int32로 문자 코드 값으로 표현
	//문자: ' '로 작성
	//자주 사용하는 escape \\, \', \", \a(콘솔벨), \b(백스페이스), \f(쪽바꿈), \r(복귀), \t(탭)
	var str1 string = "c:\\go_study\\src\\" //
	str2 := `c:\go_study\src\`

	fmt.Println(str1)
	fmt.Println(str2)

	var str3 string = "hello. world"
	var str4 string = "안녕하세요"
	var str5 string = "\ud55c\uae00"

	fmt.Println(str3)
	fmt.Println(str4)
	fmt.Println(str5)
	// fmt.Println(str2)

	//예제 3 길이
	//바이트 수를 구함
	fmt.Println(len(str3))
	fmt.Println(len(str4))
	//예제 4 길이

	fmt.Println(utf8.RuneCountInString(str3))
	fmt.Println(utf8.RuneCountInString(str4))
	fmt.Println(len([]rune(str4)))

}
