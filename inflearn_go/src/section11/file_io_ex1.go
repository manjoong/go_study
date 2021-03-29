package main

import(
	// "fmt"
	"io/ioutil"
	// "os"
)

func errCheck(e error){
	if e!= nil{
		panic(e)
	}
}

func main(){
	//파일 읽기 쓰기 -> ioutil 패키지 활용`
	//더욱 편리하고 직관적으로 파일을 읽고 쓰기 가능
	//writefile, radfile readall ..

	s:= "hELLO GO LANG"
	
	//파일 모드 -> 접근 권한
	//읽기 : 4, 쓰기 : 2, 실행 :
	//소유자, 그룹, 기타 사용자 순서 (777)

	err := ioutil.Write("test_write.txt", []byte(s). os.FileMode(0644))
	errCheck(err)


	data, err := ioutil.ReadFile("sample.txt")
	errCheck(err)

	errCheck(string(data))
	
}