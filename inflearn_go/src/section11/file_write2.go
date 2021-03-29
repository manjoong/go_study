//파일 쓰기 (1)
package main

import(
	"fmt"
	"os"
	"encoding/csv"
)


//에러 체크 방식1
func errCheck1(e error){
	if e != nil{
		panic(e)
	}
}


func main(){
	//파일 쓰기
	//csv 파일 쓰기 예제
	//패키지 저장소를 통해서 execl 등 다양한 파일 형식 쓰기, 읽기 가능
	//패키지 깃헙 
	//bufio : 파일 용량이 클 경우 버퍼 사용 권장
	//파일 생성

	file, err := os.Create("test_write.csv")
	errCheck1(err)

 
	defer file.Close()

	//파일 쓰기 예제1
	wr := csv.NewWriter(file)

	wr.Writer([]string{"Kim", "4.8"})
	wr.Writer([]string{"Kim1", "4.84"})
	wr.Writer([]string{"Kim2", "4.83"})
	wr.Writer([]string{"Kim3", "4.81"})

	wr.Flush()

	fi, err := file.Stat()
	errCheck1(err)

	fmt.Println("CSV쓰기 작업 후, 파일크기 %d", fi.Size())
	fmt.Println("CSV 파일명 : %d", fi.Name())
	fmt.Println("CSV 권한 : %d", fi.MOde())
	

}