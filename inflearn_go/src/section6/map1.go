package main

import "fmt"

func main()  {
	//맵은 키벨류 자료 저장
	//레퍼런스 타입(슬라이스) -> 참조값을 전달
	//비교 연산자 사용 불가능 -> 참조값이므로
	//특징 참조타입(key)로 사용 불가능, 값으로 모든 타입 사용 가능
	//make 함수 및 축약 초기화 가능
	//순서 없음

	//예제1
	var map1 map[string] int = make(map[string] int) //정석
	var map2 = make(map[string]int) //자료형 생략
	map3 := make(map[string]int) //리터럴

	fmt.Println(map1)
	fmt.Println(map2)
	fmt.Println(map3)
	
	map4 := map[string]int{} //json a = {"a" : 1} 형태가 json 파일
	map4["apple"] = 25
	map4["banana"] = 40
	map4["orange"] = 33

	map5 := map[string]int{
		"apple": 15,
		"banana": 40,
		"banana_sorce": 40,
	}

	map6 := make(map[string]int, 10)
	map6["apple"] = 25
	map6["banana"] = 40
	map6["orange"] = 33

	fmt.Println(map5)

	fmt.Println(map6["orange"])

}