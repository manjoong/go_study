package main

import (
	"fmt"
)

func main() {
	//문자열 표현
	//기본 utf8인코딩(유니코드 문자 집합)

	//예제 1
	var str1 string = "Golang" //
	// var str2 string = "Wolrd"  //
	var str2 string = "고프로그래밍" //

	fmt.Println(str1[0], str1[1], str1[2], str1[3], str1[4], str1[5])
	fmt.Printf("%c %c %c %c %c %c\n", str1[0], str1[1], str1[2], str1[3], str1[4], str1[5])

	fmt.Println(str2[0], str2[1], str2[2], str2[3], str2[4], str2[5])
	fmt.Printf("%c %c %c %c %c %c\n", str2[0], str2[1], str2[2], str2[3], str2[4], str2[5])

	// 예제3
	conStr := []rune(str2) //룬형으로 선언하면 깨지지 않음. 한글 정상 출력 --> 기본 패턴으로 사용
	fmt.Printf("%c %c %c %c %c %c\n", conStr[0], conStr[1], conStr[2], conStr[3], conStr[4], conStr[5])

	for i, char := range str1 {
		fmt.Printf("%c(%d)\t", char, i)
	}

	for i, char := range str2 {
		fmt.Printf("%c(%d)\t", char, i)
	}
	// fmt.Println(str2)

	// var str3 string = "hello. world"
	// var str4 string = "안녕하세요"
	// var str5 string = "\ud55c\uae00"

	// fmt.Println(str3)
	// fmt.Println(str4)
	// fmt.Println(str5)
	// // fmt.Println(str2)

	// //예제 3 길이
	// //바이트 수를 구함
	// fmt.Println(len(str3))
	// fmt.Println(len(str4))
	// //예제 4 길이

	// fmt.Println(utf8.RuneCountInString(str3))
	// fmt.Println(utf8.RuneCountInString(str4))
	// fmt.Println(len([]rune(str4)))

}
