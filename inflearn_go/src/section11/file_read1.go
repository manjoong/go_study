//파일 쓰기 (1)
package main

import(
	"fmt"
	"os"
)


//에러 체크 방식1
func errCheck1(e error){
	if e != nil{
		panic(e)
	}
}

func errCheck2(e error){
	if e != nil{
		fmt.Println(e)
		return
	}
}


func main(){
	//파일 읽기
	//open
	//close
	//read, readat
	//각 운영체제 권한 주의
	//예외 처리
	//탐생 seek 중요

	//파일 읽기 예제
	//파일 열기

	file, err := os.Open("sample.txt")
	errCheck1(err)

	//읽기 예제1
	fileInfo, err := file.Stat() //파일 사이즈 확인 위해 정보 획득
	errCheck2(err)

	fd1 := make([]byte, fileInfo.Size()) //슬라이스에 읽은 내용을 담는다
	ct1, err := file.Read(fd1)

	fmt.Println("파일 정보 출력(1)", fileInfo)
	fmt.Println("파일 이름(2)", fileInfo.Name())
	fmt.Println("파일 크기(2)", fileInfo.Size())
	fmt.Println("파일 수정 시간(2)", fileInfo.ModTime())
	fmt.Println("읽기 작업(2)", fileInfo.Name())
	fmt.Println("읽기 작업 완료 %d", ct1)
	fmt.Println(fd1)

	fmt.Println(string(fd1))

	o1, err := file.Seek(20, 0) //0은 처음 위치, 현재 위치:1,   2 마지막 위치 
	
	errCheck([]byte, 20)
	ct2, err := file.Read(fd2)

}