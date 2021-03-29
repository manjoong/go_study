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
	//파일 쓰기
	//Create : 새파일 작성 및 파일 열기
	//close 리소스 닫기
	//write, writeString, writeAt : 파일 쓰기
	//각 운영체제 권한 주의(오류 메세지 확인)
	//예외 처리 정말 중요!
	//htts://google.org/pkg/os

	//파일 쓰기 예제

	file, err := os.Create("test.txt")
	errCheck1(err)

 
	defer file.Close()

	//파일 쓰기 예제1
	s1 := []byte{7, 8, 9, 10, 11}
	n1, err := file.Write([]byte(s1))//문자열 -> byte슬라이스 형으로 변환 후 쓰기
	errCheck2(err)
	fmt.Printf("쓰기 작업 1 완료 (%d)", n1)

	file.Sync() //write commit

	//쓰기 예제2
	s2 := "hello golang firle write test! \n"
	n2, err := file.WriteString(s2)
	errCheck2(err)
	fmt.Printf("쓰기 작업2 완료 (%d)", n2)

	file.Sync() //write commit 

	s3 := "test wirte at! \n"
	n3, err := file.WriteAt([]byte(s3), 70)
	errCheck1(err)
	fmt.Printf("쓰기 작업 3 완료 (%d)", n3)

	file.Sync()

	n4, err := file.WriteString("hello golang!")
	errCheck1(err)
	fmt.Printf("쓰기 작업 4 완료 (%d)", n4)

}