//파일 쓰기 (1)
package main

import(
	"fmt"
	"os"
	"encoding/csv"
	"bufio"
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
	//csv파일 읽기 예제


	file, err := os.Open("sample.csv")
	errCheck1(err)

	defer file.Close()

	rr := csv.NewReader(bufio.NewReader(file))

	//csv 내용 읽기

	row, err := rr.Read()
	errCheck1(err)
	// row2, err2 := rr.Read()

	// errCheck1(err2)

	fmt.Println(row[0], row[1], row[7])

	fmt.Println("CSV Read Example")
	fmt.Println(row)


	// //읽기 예제1
	// fileInfo, err := file.Stat() //파일 사이즈 확인 위해 정보 획득
	// errCheck2(err)

	// fd1 := make([]byte, fileInfo.Size()) //슬라이스에 읽은 내용을 담는다
	// ct1, err := file.Read(fd1)

	rows, err := rr.ReadAll()
	errCheck2(err)
	fmt.Println(rows)

	for i, row := range rows {
		for j := range row {
			fmt.Println("%s  ", rows[i][j])
		}
	}



}
