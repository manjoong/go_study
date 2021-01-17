package main

import "fmt"

func main()  {
	//맵(map)
	//맵 조회 할 경우에 주의 할 점
	//예제1
	map1 := map[string]string{
		"다음": "www.daum.net",
		"네이버": "www.aver.com",
		"구글": "www.google.com",
		"홈": "www.test.com",
	}

	value1 := map1["다음"]
	value2 := map1["네이트온"]
	value3, ok := map1["네이트온"] //키가 존재하면 true 없으면 false


	fmt.Println(value1)
	fmt.Println(value2)
	fmt.Println(value3, ok)

	//예제2

	if value, ok := map1["네이트온"]; ok {
		fmt.Println(value)
	}else {
		fmt.Println("네이트 온은 없어")
	}

	if value, ok := map1["구글"]; ok {
		fmt.Println(value)
	}else {
		fmt.Println("구글 은 없어")
	}

	// map1["홈2"]= "www.test2.com" //추가
	// fmt.Println(map1)

	// map1["홈2"]= "www.test2-2.com" //수정
	// fmt.Println(map1)

	// delete(map1, "홈2") //삭제
	// fmt.Println(map1)

	
	// for k, v := range map1 { //맵일때는 알아서 인덱스가 아닌 키
	// 	fmt.Println(k, v)
	// }

	// for _, v := range map1 {
	// 	fmt.Println(v)
	// }
	

}