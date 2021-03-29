package main

import(
	// "fmt"
	"io/ioutil"
	"bufio"
	// "os"
)

func errCheck(e error){
	if e!= nil{
		panic(e)
	}
}

func main(){
	//파일 읽기, 버퍼사용 쓰기 -> bufio 패키지 활용
	//ioutil, bufio 등은 

	//즉 bufio 의 nw reader, new writer를 통해서 객체 반환 후 메소드 사용 가능
	//bufio

	file, err := os.OpenFile("test_write2.txt", os.O_CREATE|os.O_RDWR, os.FileMode(0777))
	wt := bufio.NewWriter(file) //writer반환

	wt.WriteString("hello golang file write test!")
	wt.Write([]byte("Hello Golang!"))
	
	//버퍼 정보 출력

	errCheck(err)

	wt.Flush() //버퍼 내용을 비루고 반영
	

	rt := bufio.NewReader(file) //reader반환
	fi, err:= file.Stat()


	file.Seek(0, os.SEEK_SET)
	data, _ := rt.Read(b)

	fmt.Printf("전체 buffer size: %d", rt.Size())

	
}