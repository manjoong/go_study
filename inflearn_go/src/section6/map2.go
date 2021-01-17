package main

import "fmt"

func main()  {
	//맵(map)
	//맵 조회 및 순회

	//예제1
	map1 := map[string]string {
		"다음": "www.daum.net",
		"네이버": "www.aver.com",
		"구글": "www.google.com",
	}

	fmt.Println(map1["구글"])

	for k, v := range map1 { //맵일때는 알아서 인덱스가 아닌 키
		fmt.Println(k, v)
	}

	for _, v := range map1 {
		fmt.Println(v)
	}
	

}